package jsonrpc

import (
	"erpcLanguageServer/jsonrpc/jsonrpcErrors"
	"errors"
	"io"
	"strconv"
	"strings"
)

type MessageReader struct {
	in io.Reader
}

/*
	Creates a new message reader
*/
func NewMessageReader(reader io.Reader) *MessageReader {
	var ret MessageReader
	ret.in = reader
	return &ret
}

/*
	Returns the next message
*/
func (reader *MessageReader) Next() ([]byte, error) {
	return reader.readRawJsonMessage()
}

/*
	Read n amount of bytes on a reader
*/
func (reader *MessageReader) readAmountOfBytes(amount int) ([]byte, error) {
	content := make([]byte, amount)
	amountRead, err := reader.in.Read(content)
	return content[:amountRead], err
}

/*
	Reads each byte of a reader until the delimeter is reached (inclusive).
	Returns all read bytes, INCLUDING the delimeter.
*/
func (reader *MessageReader) readUntil(delimeter byte) ([]byte, error) {
	var ret []byte
	for {
		r := make([]byte, 1)
		amount, err := reader.in.Read(r)
		if err != nil {
			return ret, err
		}

		if amount == 0 {
			return ret, errors.New("Unexpectedly got amount of 0 bytes while executing readUntil()")
		}

		ret = append(ret, r[0])

		if r[0] == delimeter {
			break
		}
	}
	return ret, nil
}

func (reader *MessageReader) readRawJsonMessage() ([]byte, error) {
	lengthHeader, err := reader.readAmountOfBytes(15)
	if err != nil {
		return []byte{}, err
	}

	if strings.ToLower(string(lengthHeader)) != "content-length:" {
		return []byte{}, jsonrpcErrors.ParseError{}
	}

	line, err := reader.readUntil('{')
	if err != nil {
		if err == io.EOF {
			return []byte{}, jsonrpcErrors.InternalError{Data: "Unexpected end of file in readUntil():" + err.Error()}
		} else {
			return []byte{}, jsonrpcErrors.InternalError{Data: "Could not read line:" + err.Error()}
		}
	}

	messageLength, err := strconv.Atoi(strings.TrimSpace(string(line[:len(line)-1])))
	if err != nil {
		return []byte{}, jsonrpcErrors.InternalError{Data: "Parsing the length of the message failed:" + err.Error()}
	}

	messageContent, err := reader.readAmountOfBytes(messageLength - 1) // we already read the opening {, so make it one less
	if err != nil {
		return []byte{}, err
	}

	return append([]byte{'{'}, messageContent...), err
}
