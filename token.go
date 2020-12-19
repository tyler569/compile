package main

import (
	"regexp"
	"strconv"
)

//go:generate stringer -type tokenKind
type tokenKind int

const (
	noneToken tokenKind = iota
	intToken
	stringToken
	opToken
	identToken
)

type token struct {
	kind tokenKind
	int
	string
}

type tokenStream struct {
	tokens []token
}

func (ts *tokenStream) peek() (token, bool) {
	if len(ts.tokens) == 0 {
		return token{}, false
	}
	return ts.tokens[0], true
}

func (ts *tokenStream) pop() (token, bool) {
	if len(ts.tokens) == 0 {
		return token{}, false
	}
	tok := ts.tokens[0]
	ts.tokens = ts.tokens[1:]
	// fmt.Printf("pop: %v ", tok)
	return tok, true
}

const tokenRegexString = "" +
	`'.*'|` +
	`".*"|` +
	`[0-9][0-9_]*|` +
	// `0x[0-9a-fA-F_]*|` +
	`[_a-zA-Z][_a-zA-Z0-9]*|` +
	`>>=?|` +
	`<<=?|` +
	`[+\-*/%&|^]=|` +
	`\S`

var tokenRegex = regexp.MustCompile(tokenRegexString)

func tokenSplit(s string) []string {
	return tokenRegex.FindAllString(s, -1)
}

func tokenize(s string) *tokenStream {
	split := tokenSplit(s)
	tokens := []token{}
	var err error
	for _, s := range split {
		t := token{}
		t.int, err = strconv.Atoi(s)
		if err == nil {
			t.kind = intToken
		} else {
			t.kind = opToken
			t.string = s
		}
		tokens = append(tokens, t)
	}
	return &tokenStream{tokens}
}
