package server

import (
	"encoding/json"
	"erpcLanguageServer/server/jsonrpc"
	"erpcLanguageServer/server/methods/initialize"
)

func (server *Server) OnInitialize(c func(params initialize.Parameters) (initialize.Response, *jsonrpc.JSONRPCError)) {
	server.methods[initialize.Identifier()] = func(request jsonrpc.Request) jsonrpc.Response {
		var p initialize.Parameters
		err := json.Unmarshal(request.Params, &p)
		if err != nil {
			return jsonrpc.NewInvalidParametersError("Could not parse parameters", err).ToErrorResponse(request.ID)
		}
		res, JSONRPCError := c(p)
		if JSONRPCError != nil {
			return JSONRPCError.ToErrorResponse(request.ID)
		}

		return jsonrpc.Response{
			Result: res,
			ID:     request.ID,
		}
	}
}
