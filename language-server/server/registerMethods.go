package server

import (
	"encoding/json"
	"erpcLanguageServer/server/jsonrpc"
	"erpcLanguageServer/server/methods/initialize"
	"erpcLanguageServer/server/methods/initialized"
	"erpcLanguageServer/server/methods/showMessage"
	"erpcLanguageServer/server/methods/textDocumentDidOpen"
)

/*
	Registers a callback to be executed when the initialize method gets called by the client. Overwrites the default fallback.
*/
func (server *Server) OnInitialize(c func(params initialize.Parameters) (initialize.Response, *jsonrpc.JSONRPCError)) {
	server.methods[initialize.Identifier()] = func(request jsonrpc.Request) *jsonrpc.Response {
		var p initialize.Parameters
		err := json.Unmarshal(request.Params, &p)
		if err != nil {
			return jsonrpc.NewInvalidParametersError("Could not parse parameters", err).ToErrorResponse(request.ID)
		}
		res, JSONRPCError := c(p)
		if JSONRPCError != nil {
			return JSONRPCError.ToErrorResponse(request.ID)
		}

		return &jsonrpc.Response{
			Result: res,
			ID:     request.ID,
		}
	}
}

/*
	Registers a callback to be executed when a document got opened
*/
func (server *Server) OnDidOpenTextDocument(c func(params textDocumentDidOpen.Parameters) *jsonrpc.JSONRPCError) {
	server.methods[textDocumentDidOpen.Identifier()] = func(request jsonrpc.Request) *jsonrpc.Response {
		var p textDocumentDidOpen.Parameters
		err := json.Unmarshal(request.Params, &p)
		if err != nil {
			return jsonrpc.NewInvalidParametersError("Could not parse parameters", err).ToErrorResponse(request.ID)
		}
		JSONRPCError := c(p)
		if JSONRPCError != nil {
			return JSONRPCError.ToErrorResponse(request.ID)
		}

		return &jsonrpc.Response{
			ID: request.ID,
		}
	}
}

/*
	Registers a callback to be executed when a document got opened
*/
func (server *Server) OnInitialized(c func() *jsonrpc.JSONRPCError) {
	server.methods[initialized.Identifier()] = func(request jsonrpc.Request) *jsonrpc.Response {
		JSONRPCError := c()
		if JSONRPCError != nil {
			return JSONRPCError.ToErrorResponse(request.ID)
		}

		return &jsonrpc.Response{}
	}
}

/*
	Send a message to the client. If an error occured, the client will be notified, you don't need to do this manually.
*/
func (server *Server) ShowMessage(message string, messageType showmessage.MessageType) *jsonrpc.JSONRPCError {
	return server.sendNotification(showmessage.Identifier(), showmessage.Parameters{
		MessageType: &messageType,
		Message:     &message,
	})
}
