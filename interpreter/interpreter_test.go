package interpreter

import (
	"bytes"
	"testing"
)

func TestReadName(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{
			in:  ".name: #34532627",
			out: "name",
		},
		{
			in:  ".asdgj:",
			out: "asdgj",
		},
	}

	for i := range tests {
		var name, _ = parseName([]byte(tests[i].in))
		if name != tests[i].out {
			t.Fail()
		}
	}
}

func TestReadUTF8State(t *testing.T) {
	var tests = []struct {
		in  string
		out []byte
	}{
		{
			in:  "\"34532627\"",
			out: []byte("34532627"),
		},
		{
			in:  "\"\"\"asd\"\"\"",
			out: []byte("asd"),
		},
	}

	for i := range tests {
		var state, _ = parseState([]byte(tests[i].in))
		if !bytes.Equal(state, tests[i].out) {
			t.Fail()
		}
	}
}
