package testutils

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

/*
	Reads complete file and copies it into memory, closes the file afterwards
*/
func ReadTestFile(path string, t *testing.T) io.Reader {
	file, err := os.Open(path)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log("Bytes:", len(bytes))

	file.Close()
	return strings.NewReader(string(bytes))
}
