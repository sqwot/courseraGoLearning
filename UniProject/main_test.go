package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	if err != nil {
		t.Errorf("test for OK Failed")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("test for OK failed - results not much\n %#v %#v", result, testOkResult)
	}
}

var testFaild = `1
2
1`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFaild))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	if err == nil {
		t.Errorf("Test for Failed failed")
	}
}
