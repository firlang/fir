package token

import "fmt"

const (
	PLUS  = "PLUS"
	MINUS = "MINUS"
	STAR  = "STAR"
	SLASH = "SLASH"

	COLON = "COLON"

	ASSIGN = "ASSIGN"

	NUMBER = "NUMBER"
	IDENT  = "IDENT"

	EOF = "EOF"
)

type Token struct {
	Type    string
	Literal string
	Line    int
}

func (t *Token) String() string {
	return fmt.Sprintf("{ Type: %s, Lit: %s, Line: %d }", t.Type, t.Literal, t.Line)
}
