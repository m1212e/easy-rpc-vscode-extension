package main_test

import (
	"erpcLanguageServer/jsonrpc"
	"erpcLanguageServer/testutils"
	"testing"
)



func TestReadMessage(t *testing.T) {
	file := testutils.ReadTestFile("exampleOpeningMessage", t)
	reader := jsonrpc.NewMessageReader(file)

	msg, err := reader.Next()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(msg)
}
