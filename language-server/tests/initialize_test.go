package test

import (
	"erpcLanguageServer/jsonrpc"
	"erpcLanguageServer/methods"
	"testing"
)

func TestInitialize(t *testing.T) {
	file := ReadTestFile("exampleOpeningMessage", t)
	reader := jsonrpc.NewReader(file)

	req, err := reader.Next()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	methods.Handle()
}
