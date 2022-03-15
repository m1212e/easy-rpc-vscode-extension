package server

import (
	"erpcLanguageServer/server/jsonrpc"
	"erpcLanguageServer/server/methods/initialize"
)

func (server *Server) registerDefaultMethods() {
	server.methods["exit"] = func(request jsonrpc.Request) *jsonrpc.Response {
		server.Shutdown()
		return &jsonrpc.Response{}
	}

	_, exists := server.methods[initialize.Identifier()]
	if !exists {
		server.OnInitialize(initialize.DefaultImplementation)
	}
}
