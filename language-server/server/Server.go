package server

import (
	"erpcLanguageServer/server/jsonrpc"
	"log"
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
	s.methods = make(map[string]func(request jsonrpc.Request) jsonrpc.Response)
	return s
}

func (server *Server) Run() {
	server.registerDefaultMethods()

	server.reader = jsonrpc.NewReader(os.Stdin)
	server.writer = jsonrpc.NewWriter(os.Stdout)
	server.requests = make(chan jsonrpc.Request, 5)
	server.responses = make(chan jsonrpc.Response, 5)

	server.wg.Add(3)
	go server.readerLoop()
	go server.handlerLoop()
	go server.writerLoop()
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
		log.Println("Recieved request:", request.Method, "with ID:", request.ID)
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
			log.Println("Notification recieved.")
			continue
		}

		response.Jsonrpc = "2.0"
		if err := server.writer.Write(response); err != nil {
			log.Println("Errored response with ID:", response.ID)
			server.responses <- err.ToErrorResponse(response.ID)
			continue
		}
		log.Println("Sent response with ID:", response.ID)
	}

	server.wg.Done()
}

func (server *Server) handlerLoop() {
	for {
		if server.shutdown {
			break
		}
		request := <-server.requests
		f, ok := server.methods[request.Method]
		if !ok {
			server.responses <- jsonrpc.NewMethodNotFoundError("could not find the method "+request.Method, nil).ToErrorResponse(request.ID)
			continue
		}

		server.responses <- f(request)
	}
	server.wg.Done()
}

func (server *Server) Shutdown() {
	server.shutdown = true
}
