package lexer

import (
	"fir/error"
	"fir/token"
	"fmt"
)

type Lexer struct {
	source string
	char   int
	tokens []token.Token
	line   int
}

func New(source string) Lexer {
	return Lexer{
		source: source,
		char:   0,
		line:   1,
	}
}

func (l *Lexer) Scan_tokens() []token.Token {
	c := ""
	for !l.at_end() {
		c = string(l.source[l.char])
		switch c {
		case "+":
			l.add_token(token.PLUS, "+")
		case "-":
			l.add_token(token.MINUS, "-")

		case ":":
			if l.peek_next() == "=" {
				l.add_token(token.ASSIGN, ":=")
				l.advance()
			} else {
				l.illegal(":")
			}

		case "\n":
			l.line++

		default:
			l.add_token(c, c)
		}
		l.advance()
	}
	return l.tokens
}

func (l *Lexer) advance() {
	l.char++
}

func (l *Lexer) at_end() bool {
	return l.char >= len(l.source)
}

func (l *Lexer) add_token(tok string, c string) {
	l.tokens = append(l.tokens, token.Token{
		Type:    tok,
		Literal: c,
	})
}

func (l *Lexer) peek_next() string {
	if l.char == len(l.source)-1 {
		return ""
	}
	return string(l.source[l.char+1])
}

func (l *Lexer) illegal(c string) {
	error.Error(error.LexError, fmt.Sprintf("illegal token '%s'", c), l.line)
}
