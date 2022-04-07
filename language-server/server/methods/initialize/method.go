package initialize

import (
	"erpcLanguageServer/server/jsonrpc"
	"log"
)

func Identifier() string {
	return "initialize"
}

func DefaultImplementation(params Parameters) (Response, *jsonrpc.JSONRPCError) {
	log.Println("Initialized easy-rpc language server at", *params.RootPath)
	return Response{}, nil
}
