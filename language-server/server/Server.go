package server

import (
	"erpcLanguageServer/server/jsonrpc"
	"os"
	"sync"
)

type Server struct {
	methods   map[string]func(request jsonrpc.Request) jsonrpc.Response
	reader    *jsonrpc.Reader
	writer    *jsonrpc.Writer
	requests  chan jsonrpc.Request
	responses chan jsonrpc.Response
	wg        sync.WaitGroup
	shutdown  bool
}

func NewServer() *Server {
	s := &Server{}
	return s
}

func (server *Server) Run() {
	server.registerDefaultMethods()

	server.reader = jsonrpc.NewReader(os.Stdin)
	server.writer = jsonrpc.NewWriter(os.Stdin)
	server.requests = make(chan jsonrpc.Request, 5)
	server.responses = make(chan jsonrpc.Response, 5)

	server.wg.Add(3)
	go server.readerLoop()
	go server.writerLoop()
	go server.handlerLoop()
	server.wg.Wait()
}

func (server *Server) readerLoop() {
	for {
		if server.shutdown {
			close(server.requests)
			break
		}

		request, err := server.reader.Next()
		if err != nil {
			server.responses <- err.ToErrorResponse(nil)
		}

		server.requests <- request
	}

	server.wg.Done()
}

func (server *Server) writerLoop() {
	for {
		if server.shutdown {
			close(server.responses)
			break
		}
		response := <-server.responses

		// dont send notification answers
		if response.ID == nil {
			continue
		}

		response.Jsonrpc = "2.0"

		if err := server.writer.Write(response); err != nil {
			server.responses <- err.ToErrorResponse(nil)
		}
	}

	server.wg.Done()
}

func (server *Server) handlerLoop() {
	for {
		if server.shutdown {
			break
		}
		request := <-server.requests
		server.responses <- server.handleIncomingMethod(request)
	}
	server.wg.Done()
}

func (server *Server) handleIncomingMethod(request jsonrpc.Request) jsonrpc.Response {
	f, ok := server.methods[request.Method]
	if !ok {
		return jsonrpc.NewMethodNotFoundError("could not find the method "+request.Method, nil).ToErrorResponse(request.ID)
	}

	return f(request)
}

func (server *Server) Shutdown() {
	server.shutdown = true
}
