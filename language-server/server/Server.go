package server

import (
	"crypto/rand"
	"encoding/json"
	"erpcLanguageServer/server/jsonrpc"
	"log"
	"os"
	"sync"
	"time"
)

type Server struct {
	methods         map[string]func(request jsonrpc.Request) *jsonrpc.Response
	ongoingRequests map[string]chan jsonrpc.Response
	reader          *jsonrpc.Reader
	writer          *jsonrpc.Writer
	incoming        chan jsonrpc.Sendable
	outgoing        chan jsonrpc.Sendable
	wg              sync.WaitGroup
	shutdown        bool
}

func NewServer() *Server {
	s := &Server{}
	s.methods = make(map[string]func(request jsonrpc.Request) *jsonrpc.Response)
	s.ongoingRequests = make(map[string]chan jsonrpc.Response)
	return s
}

func (server *Server) Run() {
	server.registerDefaultMethods()

	server.reader = jsonrpc.NewReader(os.Stdin)
	server.writer = jsonrpc.NewWriter(os.Stdout)
	server.incoming = make(chan jsonrpc.Sendable, 5)
	server.outgoing = make(chan jsonrpc.Sendable, 5)

	server.wg.Add(3)
	go server.readerLoop()
	go server.handlerLoop()
	go server.writerLoop()
	server.wg.Wait()
}

func (server *Server) readerLoop() {
	for {
		if server.shutdown {
			close(server.incoming)
			break
		}

		message, err := server.reader.Next()
		if err != nil {
			server.outgoing <- err.ToErrorResponse(nil)
		}
		log.Println("Recieved incoming:", message.SendableToString())
		server.incoming <- message
	}

	server.wg.Done()
}

func (server *Server) writerLoop() {
	for {
		if server.shutdown {
			close(server.outgoing)
			break
		}
		message := <-server.outgoing

		// dont send notification responses
		if response, ok := message.(*jsonrpc.Response); ok && response.ID == nil {
			continue
		}

		if err := server.writer.Write(message); err != nil {
			log.Println("Outgoing message errored:" + message.SendableToString())
			server.outgoing <- err.ToErrorResponse(message.GetID())
			continue
		}
		log.Println("Sent message with ID:", message.GetID())
	}

	server.wg.Done()
}

//TODO refactor

func (server *Server) handlerLoop() {
	for {
		if server.shutdown {
			break
		}
		message := <-server.incoming

		go func() {
			if req, ok := message.(*jsonrpc.Request); ok {
				f, ok := server.methods[req.Method]
				if !ok {
					log.Println("Could not find method:", req.Method)
					server.outgoing <- jsonrpc.NewMethodNotFoundError("could not find the method "+req.Method, nil).ToErrorResponse(req.ID)
					return
				}

				server.outgoing <- f(*req)
				return
			}

			if res, ok := message.(*jsonrpc.Response); ok {
				if id, ok := res.ID.(string); ok {
					_, ok := server.ongoingRequests[id]
					if ok {
						server.ongoingRequests[id] <- *res
					} else {
						server.outgoing <- jsonrpc.NewInternalError("id not found, request might have timed out", nil).ToErrorResponse(res.ID)
					}
					return
				}

				server.outgoing <- jsonrpc.NewInternalError("invalid id for response, must be string type", nil).ToErrorResponse(res.ID)
				return
			}

			server.outgoing <- jsonrpc.NewInternalError("invalid state, message must be request or response", nil).ToErrorResponse(nil)
		}()
	}
	server.wg.Done()
}

/*
	Sends a request to the server. Returns an error which may have occured. You do not need to send that error to the client,
	this is done automatically. Times out after 5 seconds.
*/
func (server *Server) makeRequest(method string, params interface{}) (*jsonrpc.Response, *jsonrpc.JSONRPCError) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		e := jsonrpc.NewInternalError("Could not generate a request ID", err)
		server.outgoing <- e.ToErrorResponse(nil)
		return nil, e
	}
	id := string(b)

	resolve := make(chan jsonrpc.Response)
	server.ongoingRequests[id] = resolve
	defer delete(server.ongoingRequests, id)

	data, err := json.Marshal(params)
	if err != nil {
		e := jsonrpc.NewInternalError("Could not marshal parameters", err)
		server.outgoing <- e.ToErrorResponse(nil)
		return nil, e
	}

	server.outgoing <- &jsonrpc.Request{
		Method: method,
		Params: data,
		ID:     id,
	}

	select {
	case res := <-resolve:
		return &res, nil
	case <-time.After(5 * time.Second):
		e := jsonrpc.NewInternalError("Request with id "+id+" timed out", nil)
		server.outgoing <- e.ToErrorResponse(nil)
		return nil, e
	}
}

/*
	Sends a notification to the server. Returns an error which may have occured. You do not need to send that error to the client,
	this is done automatically.
*/
func (server *Server) sendNotification(method string, params interface{}) *jsonrpc.JSONRPCError {
	data, err := json.Marshal(params)
	if err != nil {
		e := jsonrpc.NewInternalError("Could not marshal parameters", err)
		server.outgoing <- e.ToErrorResponse(nil)
		return e
	}

	server.outgoing <- &jsonrpc.Request{
		Method: method,
		Params: data,
	}
	return nil
}

func (server *Server) Shutdown() {
	server.shutdown = true
}
