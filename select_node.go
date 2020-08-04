package torm

import "strings"

type SelectNode struct {
	nt      NodeType
	columns []string
	next    Node
}

func (n *SelectNode) Next() (next Node, err error) {
	return
}

func (n *SelectNode) Walk(build strings.Builder) (err error) {
	return
}
