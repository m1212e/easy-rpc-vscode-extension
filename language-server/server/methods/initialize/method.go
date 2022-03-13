package initialize

import "erpcLanguageServer/server/jsonrpc"

func Identifier() string {
	return "initialize"
}

func DefaultImplementation(params Parameters) (Response, *jsonrpc.JSONRPCError) {
	return Response{
		Capabilities: ServerCapabilities{},
	}, nil
}
