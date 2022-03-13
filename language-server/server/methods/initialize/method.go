package initialize

import "erpcLanguageServer/server/jsonrpc"

func Identifier() string  {
	return "initialize"
}

func DefaultImplementation(params Params) (Response, *jsonrpc.JSONRPCError) {
	return Response{
		
	}, nil
}
