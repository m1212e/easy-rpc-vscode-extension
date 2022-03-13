package jsonrpc

import "fmt"

/*
	Error which can occur in the application
*/
type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (err JSONRPCError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, data: %v", err.Code, err.Message, err.Data)
}

/*
	Converts this error to an error response which can be sent to the client
*/
func (err JSONRPCError) ToErrorResponse(ID interface{}) Response {
	return Response{
		Result: nil,
		ID:     ID,
		Error:  &err,
	}
}

/*
	Internal JSON-RPC error.
*/
func NewInternalError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32603,
		Message: message,
		Data:    data,
	}
}

/*
	Invalid method parameter(s).
*/
func NewInvalidParametersError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32602,
		Message: message,
		Data:    data,
	}
}

/*
	The JSON sent is not a valid Request object.
*/
func NewInvalidRequestError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32600,
		Message: message,
		Data:    data,
	}
}

/*
	The method does not exist / is not available.
*/
func NewMethodNotFoundError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32601,
		Message: message,
		Data:    data,
	}
}

/*
	Invalid JSON was received by the server.
	An error occurred on the server while parsing the JSON text.
*/
func NewParseError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32700,
		Message: message,
		Data:    data,
	}
}

/*
	Error code indicating that a server received a notification or
	request before the server has received the `initialize` request.
*/
func NewServerNotInitializedError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32002,
		Message: message,
		Data:    data,
	}
}

/*
	Indicates an unknown error code
*/
func NewUnknownErrorCodeError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32001,
		Message: message,
		Data:    data,
	}
}

/*
	Indicates that the content has been modified
*/
func NewContentModifiedError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32801,
		Message: message,
		Data:    data,
	}
}

/*
	Indicates that the content has been modified
*/
func NewRequestCancelledError(message string, data interface{}) *JSONRPCError {
	return &JSONRPCError{
		Code:    -32800,
		Message: message,
		Data:    data,
	}
}
