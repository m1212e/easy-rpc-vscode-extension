package test

import (
	"erpcLanguageServer/jsonrpc"
	"testing"
)

func TestReadMessage(t *testing.T) {
	file := ReadTestFile("exampleOpeningMessage", t)
	reader := jsonrpc.NewReader(file)

	msg, err := reader.Next()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	AssertEquals(msg.Jsonrpc, "2.0", t)
	AssertEquals(msg.ID, float64(0), t)
	AssertEquals(msg.Method, "initialize", t)
}
