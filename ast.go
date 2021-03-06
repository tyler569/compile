package main

import (
	"fmt"
)

//go:generate stringer -type nodeKind
type nodeKind int

const (
	number nodeKind = iota
	str
	ident
	binop
)

type node struct {
	kind   nodeKind
	int           // number
	string        // ident, str
	op     string // binary operator

	// binop
	left  *node
	right *node
}

func (n *node) String() string {
	switch n.kind {
	case number:
		return fmt.Sprint(n.int)
	case str:
		return fmt.Sprintf("\"%v\"", n.string)
	case ident:
		return fmt.Sprint(n.string)
	case binop:
		return fmt.Sprintf("(%v %v %v)", n.left, n.op, n.right)
	default:
		fatalf("Invalid node\n")
		return ""
	}
}
