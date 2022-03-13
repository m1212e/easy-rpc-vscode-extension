package main

import (
	"erpcLanguageServer/server"
	"erpcLanguageServer/server/jsonrpc"
	"erpcLanguageServer/server/methods/initialize"
)

func main() {
	server := server.NewServer()

	server.OnInitialize(func(params initialize.Parameters) (initialize.Response, *jsonrpc.JSONRPCError) {
		//TODO start server
		// params.RootPath
		return initialize.DefaultImplementation(params)
	})

	server.Run()
}
