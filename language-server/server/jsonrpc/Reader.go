package jsonrpc

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type Reader struct {
	in io.Reader
}

/*
	Creates a new message reader
*/
func NewReader(reader io.Reader) *Reader {
	var ret Reader
	ret.in = reader
	return &ret
}

/*
	Returns the next message
*/
func (reader *Reader) Next() (Sendable, *JSONRPCError) {
	raw, err := reader.readRawJsonMessage()
	if err != nil {
		return nil, err
	}

	var request Request
	var response Response
	jsonErrorRequest := json.Unmarshal(raw, &request)
	jsonErrorResponse := json.Unmarshal(raw, &response)

	if jsonErrorRequest != nil && jsonErrorResponse != nil {
		return nil, NewInternalError("error while unmarshalling json message", jsonErrorRequest)
	}

	if jsonErrorRequest != nil {
		return &response, nil
	} else {
		return &request, nil
	}
}

/*
	Read n amount of bytes on a reader
*/
func (reader *Reader) readAmountOfBytes(amount int) ([]byte, *JSONRPCError) {
	content := make([]byte, amount)
	amountRead, err := reader.in.Read(content)

	if err != nil {
		return content[:amountRead], NewInternalError("Error while reading bytes:", err.Error())
	}
	return content[:amountRead], nil
}

/*
	Reads each byte of a reader until the delimeter is reached (inclusive).
	Returns all read bytes, INCLUDING the delimeter.
*/
func (reader *Reader) readUntil(delimeter byte) ([]byte, *JSONRPCError) {
	var ret []byte
	for {
		r := make([]byte, 1)
		amount, err := reader.in.Read(r)
		if err != nil {
			return ret, NewInternalError("could not read next byte in readUntil()", err)
		}

		if amount == 0 {
			return ret, NewInternalError("unexpectedly got amount of 0 bytes while executing readUntil()", nil)
		}

		ret = append(ret, r[0])

		if r[0] == delimeter {
			break
		}
	}
	return ret, nil
}

func (reader *Reader) readRawJsonMessage() ([]byte, *JSONRPCError) {
	lengthHeader, err := reader.readAmountOfBytes(15)
	if err != nil {
		return []byte{}, err
	}

	if strings.ToLower(string(lengthHeader)) != "content-length:" {
		return []byte{}, NewParseError("could not parse content length, found string didnt match", "found:"+string(lengthHeader))
	}

	//TODO: Handle Content-Type header part accordingly
	line, err := reader.readUntil('{')
	if err != nil {
		if err == io.EOF {
			return []byte{}, NewParseError("unexpected end of file in readUntil()", err.Error())
		} else {
			return []byte{}, NewParseError("could not read line", err.Error())
		}
	}

	messageLength, atoiError := strconv.Atoi(strings.TrimSpace(string(line[:len(line)-1])))
	if atoiError != nil {
		return []byte{}, NewParseError("parsing the length of the message JSONRPCErrored", err.Error())
	}

	messageContent, err := reader.readAmountOfBytes(messageLength - 1) // we already read the opening { in readUntil(), so make it one less
	if err != nil {
		return []byte{}, err
	}

	return append([]byte{'{'}, messageContent...), err
}
