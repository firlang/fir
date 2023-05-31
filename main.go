package main

import (
	"fir/lexer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		run_file(os.Args[1])
	}
}

func run_file(source string) {
	f, err := os.ReadFile(source)

	if err != nil {
		panic(err)
	}

	l := lexer.New(string(f))

	for _, tok := range l.Scan_tokens() {
		fmt.Println(tok.String())
	}
}
