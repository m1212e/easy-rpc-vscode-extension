package jsonrpc

import (
	"encoding/json"
	"fmt"
)

type Sendable interface {
	SendableToString() string
	SetJSONRPCVersion(string)
	GetID() interface{}
}

/*
	A request
*/
type Request struct {
	Jsonrpc string          `json:"jsonrpc,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      interface{}     `json:"id,omitempty"`
}

func (req Request) SendableToString() string {
	return fmt.Sprintf("{jsonrpc: %v, method: %v, params: %+v, ID: %v}", req.Jsonrpc, req.Method, req.Params, req.ID)
}
func (req Request) GetID() interface{} {
	return req.ID
}
func (req *Request) SetJSONRPCVersion(v string) {
	req.Jsonrpc = v
}

/*
	A response
*/
type Response struct {
	Jsonrpc string        `json:"jsonrpc,omitempty"`
	Result  interface{}   `json:"result,omitempty"`
	ID      interface{}   `json:"id,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

func (res Response) SendableToString() string {
	return fmt.Sprintf("{jsonrpc: %v, result: %v, error: %+v, ID: %v}", res.Jsonrpc, res.Result, res.Error, res.ID)
}
func (res *Response) SetJSONRPCVersion(v string) {
	res.Jsonrpc = v
}
func (res Response) GetID() interface{} {
	return res.ID
}
