package interpreter

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/ki/interpreter/tree"
	"io"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

const (
	op_name_start   = '.'
	op_name_end     = ':'
	op_state        = '#'
	op_state_string = '"'
	op_brunch_start = '{'
	op_brunch_end   = '}'
)

type Worker interface {
	Parse() error
	Interpret() []string
}

type job struct {
	data io.Reader
	tree *tree.Node
}

func parseName(data []byte) (string, int) {
	var end = strings.IndexRune(string(data), op_name_end)
	return string(data[:end]), end
}

func parseState(data []byte) ([]byte, int) {
	var start int

	for i, c := range string(data) {
		if c == op_state_string {
			start = i + 1
		} else {
			break
		}
	}

	if start == 0 {
		return nil, 0
	}

	var end = strings.Index(string(data[start:]), strings.Repeat(string(op_state_string), start)) + start
	data = data[start:end]
	return data, end
}

func (j *job) parse(data []byte) error {
	var t = j.tree
	var i, w int
	var r rune

	var name string
	var ready = true

	for i < len(data) {
		r, w = utf8.DecodeRune(data[i:])
		i += w

		switch r {
		case op_name_start:
			if !ready {
				return erres.Error("not ready")
			}

			var n, ii = parseName(data[i:])
			i += ii
			name = n
			ready = false
		case op_state:
			if ready {
				return erres.Error("ready")
			}

			var state, ii = parseState(data[i:])
			i += ii

			t.AddStateNode(name, state)
			ready = true
		case op_brunch_start:
			if ready {
				return erres.Error("ready")
			}

			t = t.AddBranchNode(name)
			ready = true
		case op_brunch_end:
			if !ready {
				return erres.Error("not ready")
			}

			t = t.GetParent()
		}

	}

	return nil
}

func (j *job) Parse() error {
	var data, err = ioutil.ReadAll(j.data)
	if err != nil {
		return err
	}

	return j.parse(data)
}

func (j *job) interpret() []string {
	var indexes = make([]int, 256)
	var re []string

dfs:
	for {
		var names []string
		var state []byte
		var t = j.tree
		var depth int
		for {
			names = append(names, t.GetName())
			if t.IsLeaf() {
				state = t.GetState()
				indexes[depth]++
				depth = 0
				t = j.tree
				break
			} else {
				depth++
				var tt = t.GetChild(indexes[depth])
				if tt != nil {
					t = tt
				} else {
					if depth == 1 {
						break dfs
					} else {
						indexes[depth] = 0
						indexes[depth-1]++
						t = j.tree
						depth = 0
						names = nil
						break
					}
				}
			}
		}

		if names != nil {
			var name = strings.Join(names, ".")
			var line = strings.Join([]string{name, string(state)}, ": ")
			re = append(re, line)
		}
	}

	return re
}

func (j *job) Interpret() []string {
	return j.interpret()
}

func NewJob(src io.Reader) (*job, error) {
	if src == nil {
		return nil, erres.NilArgument
	}

	var j = new(job)
	j.data = src

	var err error
	if j.tree, err = tree.NewTree("main"); err != nil {
		return nil, err
	}

	return j, nil
}
