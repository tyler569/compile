package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func eval(n *node) int {
	if n.kind == number {
		return n.int
	}
	if n.kind != binop {
		fatalf("Can't eval %v node\n", n.kind)
	}
	if n.left == n {
		fatalf("cycle detected\n")
	}
	l := eval(n.left)
	r := eval(n.right)

	if *verbose {
		fmt.Printf("%v %v %v\n", l, "+", r)
	}

	switch n.op {
	case "+":
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	case "/":
		return l / r
	}
	fatalf("invalid eval")
	return 0
}

var file = flag.String("file", "test_input", "Input file")
var expr = flag.String("expr", "", "Raw expression to interpret")
var verbose = flag.Bool("verbose", false, "Be verbose")

func evalLine(parseFn func(*tokenStream) *node, line string) int {
	tokens := tokenize(line)
	if len(tokens.tokens) == 0 {
		return 0
	}
	if *verbose {
		fmt.Println(tokens.tokens)
	}
	ast := parseFn(tokens)
	result := eval(ast)
	if *verbose {
		fmt.Println("result:", result)
	}
	return result
}

func main() {
	flag.Parse()
	var content string

	if *expr != "" {
		content = *expr
	} else {
		bcontent, err := ioutil.ReadFile(*file)
		if err != nil {
			fatalf("%v\n", err)
		}
		content = string(bcontent)
	}

	ts := tokenize(content)
	for t, ok := ts.pop(); ok; t, ok = ts.pop() {
		fmt.Printf("%+v\n", t)
	}

	ts = tokenize(content)
	ast := parse2Sum(ts)
	fmt.Printf("ast: %v\n", ast)

	result := eval(ast)
	fmt.Println("result:", result)
}

func fatalf(format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, format, values...)
	os.Exit(1)
}
