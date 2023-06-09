package lexer

import (
	"fir/error"
	"fir/token"
	"fmt"
	"strings"
)

type Lexer struct {
	source    string
	char      int
	tokens    []token.Token
	line      int
	had_error bool
}

func New(source string) Lexer {
	return Lexer{
		source:    source,
		char:      0,
		line:      1,
		had_error: false,
	}
}

func (l *Lexer) Scan_tokens() ([]token.Token, bool) {
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
		case "%":
			l.add_token(token.MODULO, "%")
		case "^":
			l.add_token(token.CARET, "^")
		case ".":
			l.add_token(token.DOT, ".")
		case "=":
			if l.peek_next() == "=" {
				l.add_token(token.EQUALTO, "==")
				l.advance()
			}
		case ">":
			if l.peek_next() == "=" {
				l.add_token(token.GTEQUALTO, ">=")
				l.advance()
			}
		case "<":
			if l.peek_next() == "=" {
				l.add_token(token.LTEQUALTO, "<=")
				l.advance()
			}
		case "&":
			if l.peek_next() == "&" {
				l.add_token(token.AND, "&&")
				l.advance()
			}

		case "\"":
			l.string()

		case "(":
			l.add_token(token.LPAREN, "(")
		case ")":
			l.add_token(token.RPAREN, ")")
		case "{":
			l.add_token(token.LBRACE, "{")
		case "}":
			l.add_token(token.RBRACE, "}")

		case ";":
			l.add_token(token.SEMICOLON, ";")

		case ":":
			if l.peek_next() == "=" {
				l.add_token(token.ASSIGN, ":=")
				l.advance()
			} else {
				l.illegal(":")
			}
		case ",":
			l.add_token(token.COMMA, ",")

		case "\n":
			l.line++

		default:
			if l.is_digit(c) {
				l.number()
			} else if l.is_alpha(c) {
				l.ident()
			} else {
				if c != " " {
					l.illegal(c)
				}
			}
		}
		l.advance()
	}
	l.add_token(token.EOF, "")
	return l.tokens, l.had_error
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
	error.Error(error.LexError, fmt.Sprintf("illegal token '%s'", c), l.line, strings.Split(l.source, "\n")[l.line-1])
	l.had_error = true
}

func (l *Lexer) is_digit(c string) bool {
	return c >= "0" && c <= "9"
}

func (l *Lexer) is_alpha(c string) bool {
	return (c >= "a" && c <= "z") ||
		(c >= "A" && c <= "Z") ||
		c == "_"
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

	if l.peek_next() == "." {
		buf = append(buf, ".")
		l.advance()
		l.advance()
		for l.is_digit(l.peek()) {
			buf = append(buf, string(l.source[l.char]))
			if l.is_digit(l.peek_next()) {
				l.advance()
			} else {
				break
			}
		}
	}

	l.add_token(token.NUMBER, strings.Join(buf, ""))
}

func (l *Lexer) string() {
	buf := []string{}

	l.advance()

	for {
		if !l.at_end() {
			if l.peek() == "\"" {
				break
			}
			buf = append(buf, string(l.source[l.char]))
			l.advance()

		} else {
			error.Error(error.LexError, "unterminated string", l.line, strings.Split(l.source, "\n")[l.line-1])
			l.had_error = true
			break
		}
	}

	l.add_token(token.STRING, strings.Join(buf, ""))
}

func (l *Lexer) ident() {
	buf := []string{}

	for l.is_alpha(l.peek()) {
		buf = append(buf, string(l.source[l.char]))
		if l.is_alpha(l.peek_next()) {
			l.advance()
		} else {
			break
		}
	}

	ident := strings.Join(buf, "")
	switch ident {
	case "true":
		l.add_token(token.TRUE, "true")
	case "false":
		l.add_token(token.FALSE, "false")
	case "fn":
		l.add_token(token.FN, "fn")
	case "if":
		l.add_token(token.IF, "if")
	default:
		l.add_token(token.IDENT, ident)
	}
}
