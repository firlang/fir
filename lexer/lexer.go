package lexer

import (
	"fir/error"
	"fir/token"
	"fmt"
	"strings"
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
		case "*":
			l.add_token(token.STAR, "*")
		case "/":
			l.add_token(token.SLASH, "/")

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
			if l.is_digit(c) {
				l.number()
			} else {
				if c != " " {
					l.illegal(c)
				}
			}
		}
		l.advance()
	}
	l.add_token(token.EOF, "")
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
		Line:    l.line,
	})
}

func (l *Lexer) peek_next() string {
	if l.char == len(l.source)-1 {
		return ""
	}
	return string(l.source[l.char+1])
}

func (l *Lexer) peek() string {
	if l.char == len(l.source) {
		return ""
	}
	return string(l.source[l.char])
}

func (l *Lexer) illegal(c string) {
	error.Error(error.LexError, fmt.Sprintf("illegal token '%s'", c), l.line)
}

func (l *Lexer) is_digit(c string) bool {
	return c >= "0" && c <= "9"
}

func (l *Lexer) number() {
	buf := []string{}

	for l.is_digit(l.peek()) {
		buf = append(buf, string(l.source[l.char]))
		if l.is_digit(l.peek_next()) {
			l.advance()
		} else {
			break
		}
	}

	l.add_token(token.NUMBER, strings.Join(buf, ""))
}
