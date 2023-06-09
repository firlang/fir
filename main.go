package main

import (
	"fir/lexer"
	"fir/repl"
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		repl.Repl()
	case 2:
		run_file(os.Args[1])
	default:
		println("usage: fir <file>")
	}
}

func run_file(source string) {
	f, err := os.ReadFile(source)

	if err != nil {
		panic(err)
	}

	l := lexer.New(string(f))
	tokens, e := l.Scan_tokens()
	if e {
		os.Exit(1)
	}
	for _, tok := range tokens {
		fmt.Println(tok.String())
	}
}
