package token

import "fmt"

const (
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	STAR      = "STAR"
	SLASH     = "SLASH"
	MODULO    = "MODULO"
	CARET     = "CARET"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"
	SEMICOLON = "SEMICOLON"

	COLON = "COLON"
	DOT   = "DOT"
	COMMA = "COMMA"

	ASSIGN    = "ASSIGN"
	EQUALTO   = "EQUALTO"
	GTEQUALTO = "GTEQUALTO"
	LTEQUALTO = "LTEQUALTO"
	AND       = "AND"

	NUMBER = "NUMBER"
	IDENT  = "IDENT"
	STRING = "STRING"

	TRUE  = "TRUE"
	FALSE = "FALSE"
	FN    = "FN"
	IF    = "IF"

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
