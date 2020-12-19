package main

import (
	"fmt"
)

func parse2Int(tokens *tokenStream) *node {
	t, ok := tokens.pop()
	if !ok || t.kind != intToken {
		fatalf("expected int, got %v\n", t)
	}

	return &node{kind: number, int: t.int}
}

func parse2Paren(tokens *tokenStream) *node {
	if *verbose {
		fmt.Println("parse2Paren:", tokens.tokens)
	}

	t, ok := tokens.peek()
	if ok && t.kind == intToken {
		return parse2Int(tokens)
	}
	if !ok || t.string != "(" {
		fatalf("expected paren, got ? (%v)\n", tokens.tokens)
	}

	tokens.pop()
	factor := parse2Sum(tokens)

	t, ok = tokens.peek()
	if ok && t.string == ")" {
		tokens.pop()
	}

	return factor
}

func samePrecedence(t token, s string) bool {
	if t.kind != opToken {
		return false
	}

	if s == "+" && (t.string == "+" || t.string == "-") {
		return true
	}
	if s == "*" && (t.string == "*" || t.string == "/") {
		return true
	}
	return false
}

func parse2Sum(tokens *tokenStream) *node {
	if *verbose {
		fmt.Println("parse2Sum:", tokens.tokens)
	}

	left := parse2Factor(tokens)
	op, ok := tokens.peek()
	if !ok || !samePrecedence(op, "+") {
		return left
	}
	tokens.pop()
	right := parse2Factor(tokens)
	n := &node{
		left:  left,
		right: right,
		kind:  binop,
		op:    op.string,
	}

	for {
		op, ok = tokens.peek()
		if !ok || !samePrecedence(op, "+") {
			return n
		}
		tokens.pop()

		rhs := parse2Factor(tokens)
		n = &node{
			left:  n,
			right: rhs,
			kind:  binop,
			op:    op.string,
		}
	}
}

func parse2Factor(tokens *tokenStream) *node {
	if *verbose {
		fmt.Println("parse2Factor:", tokens.tokens)
	}

	left := parse2Paren(tokens)
	op, ok := tokens.peek()
	if !ok || !samePrecedence(op, "*") {
		return left
	}
	tokens.pop()
	right := parse2Paren(tokens)
	n := &node{
		left:  left,
		right: right,
		kind:  binop,
		op:    op.string,
	}

	for {
		op, ok = tokens.peek()
		if !ok || !samePrecedence(op, "*") {
			return n
		}
		tokens.pop()

		rhs := parse2Paren(tokens)
		n = &node{
			left:  n,
			right: rhs,
			kind:  binop,
			op:    op.string,
		}
	}
}
