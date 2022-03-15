package main

import (
	"erpcLanguageServer/server"
	"erpcLanguageServer/server/jsonrpc"
	"erpcLanguageServer/server/methods/initialize"
	"erpcLanguageServer/server/methods/showMessage"
)

func main() {
	server := server.NewServer()
	// rootPath := ""
	server.OnInitialize(func(params initialize.Parameters) (initialize.Response, *jsonrpc.JSONRPCError) {
		//TODO start compiler
		// rootPath = *params.RootPath
		return initialize.DefaultImplementation(params)
	})

	server.OnInitialized(func() *jsonrpc.JSONRPCError {
		server.ShowMessage("Easy-RPC language server initialized successfully", showmessage.Info)
		return nil
	})

	server.Run()
}
