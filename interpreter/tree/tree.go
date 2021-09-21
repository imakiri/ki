package tree

import "github.com/imakiri/erres"

type Node struct {
	parent   *Node
	children []*Node
	name     string
	state    []byte
}

func (n *Node) GetName() string {
	return n.name
}

func (n *Node) GetState() []byte {
	return n.state
}

func (n *Node) IsLeaf() bool {
	if n.children == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) GetChild(i int) *Node {
	if 0 <= i && i < len(n.children) {
		return n.children[i]
	} else {
		return nil
	}
}

func (n *Node) SetState(state []byte) {
	n.state = state
}

func (n *Node) AddStateNode(name string, state []byte) {
	var no = newNode(name, n)
	no.state = state
	n.children = append(n.children, no)
}

func (n *Node) AddBranchNode(name string) *Node {
	var no = newNode(name, n)
	n.children = append(n.children, no)
	return no
}

func (n *Node) GetParent() *Node {
	return n.parent
}

func newNode(name string, parent *Node) *Node {
	var n = new(Node)
	n.name = name
	n.parent = parent
	return n
}

func NewTree(name string) (*Node, error) {
	var root = new(Node)

	if name == "" {
		return nil, erres.NilArgument
	}

	root.name = name
	return root, nil
}
