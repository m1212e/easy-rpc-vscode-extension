package jsonrpc

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type Writer struct {
	out io.Writer
}

/*
	Creates a new message reader
*/
func NewWriter(writer io.Writer) *Writer {
	var ret Writer
	ret.out = writer
	return &ret
}

func (writer *Writer) Write(message Sendable) *JSONRPCError {
	message.SetJSONRPCVersion("2.0")
	bytes, err := json.Marshal(message)
	if err != nil {
		return NewInternalError("could not marshal outgoing", err)
	}

	prefix := []byte(fmt.Sprintf("Content-Length: %d\r\n\r\n", len(bytes)))

	bytes = append(prefix, bytes...)

	amountWritten, err := writer.out.Write(bytes)
	if err != nil {
		return NewInternalError("error while writing to stdout", err)
	}
	if amountWritten != len(bytes) {
		return NewInternalError("not all bytes could be written to stdout, written: "+strconv.FormatInt(int64(amountWritten), 10)+" total: "+strconv.FormatInt(int64(len(bytes)), 10), nil)
	}

	return nil
}
