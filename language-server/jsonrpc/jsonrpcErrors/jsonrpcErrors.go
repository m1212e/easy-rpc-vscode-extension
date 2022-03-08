package jsonrpcErrors

import "fmt"

type ParseError struct {
	Data interface{}
}

func (err ParseError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, data: %v", -32700, "Parse error", err.Data)
}

type InvalidRequest struct {
	Data interface{}
}

func (err InvalidRequest) Error() string {
	return fmt.Sprintf("code: %d, message: %s, data: %v", -32600, "Invalid Request", err.Data)
}

type MethodNotFound struct {
	Data interface{}
}

func (err MethodNotFound) Error() string {
	return fmt.Sprintf("code: %d, message: %s, data: %v", -32601, "Method not found", err.Data)
}

type InvalidParameters struct {
	Data interface{}
}

func (err InvalidParameters) Error() string {
	return fmt.Sprintf("code: %d, message: %s, data: %v", -32602, "Invalid params", err.Data)
}

type InternalError struct {
	Data interface{}
}

func (err InternalError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, data: %v", -32603, "Internal error", err.Data)
}
