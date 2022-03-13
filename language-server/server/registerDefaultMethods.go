package server

import "erpcLanguageServer/server/methods/initialize"

func (server *Server) registerDefaultMethods() {
	_, exists := server.methods[initialize.Identifier()]
	if !exists {
		server.OnInitialize(initialize.DefaultImplementation)
	}
}
