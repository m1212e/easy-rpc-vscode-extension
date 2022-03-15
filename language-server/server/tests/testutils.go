package test

import (
	"io"
	"io/ioutil"
	"os"
	"reflect"
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

	file.Close()
	return strings.NewReader(string(bytes))
}

/*
	Checks equality of two objects via deep equals and fails the test if they dont match
*/
func AssertEquals(got, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(got, expected) {
		t.Log("Expected", expected, "instead of", got)
		t.Fail()
	}
}
