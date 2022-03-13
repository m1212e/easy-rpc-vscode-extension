package jsonrpc

import (
	"encoding/json"
)

/*
	An incoming request
*/
type Request struct {
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      interface{}     `json:"id,omitempty"`
}

/*
	A response
*/
type Response struct {
	Jsonrpc string       `json:"jsonrpc"`
	Result  interface{}  `json:"result,omitempty"`
	ID      interface{}  `json:"id,omitempty"`
	Error   *JSONRPCError `json:"error"`
}
