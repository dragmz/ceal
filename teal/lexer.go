package teal

import (
	"strings"
	"unicode/utf8"
)

type TokenType string

const (
	TokenEol       = "EOL"
	TokenSemicolon = ";"

	TokenValue = "Value" // value

	TokenComment = "Comment" // value
)

type Token struct {
	v string // value

	l int // line
	b int // begin
	e int // end

	t TokenType
}

func (t Token) IsSeparator() bool {
	switch t.t {
	case TokenEol:
		return true
	case TokenSemicolon:
		return true
	}

	return false
}

func (t Token) StartLine() int {
	return t.l
}

func (t Token) StartCharacter() int {
	return t.b
}

func (t Token) EndLine() int {
	return t.l
}

func (t Token) EndCharacter() int {
	return t.e
}

func (t Token) String() string {
	return t.v
}

func (t Token) Line() int {
	return t.l
}

func (t Token) Begin() int {
	return t.b
}

func (t Token) End() int {
	return t.e
}

func (t Token) Type() TokenType {
	return t.t
}

type Lexer struct {
	p int // previous index
	i int // current index

	l  int // current line index
	lb int // current line begin index in source
	li int // current index in the current line

	ts  []Token // tokens
	tsi int     // tokens index

	diag []lexerError // errors

	Source []byte
}

func (z *Lexer) fail(msg string) {
	z.diag = append(z.diag, lexerError{
		l:  z.l,
		li: z.li,

		msg: msg,
	})
}

func (z *Lexer) inc(n int) {
	z.i += n
	z.li += n
}

func (z *Lexer) emit(t TokenType) {
	z.ts = append(z.ts, Token{
		l: z.l,
		b: z.p - z.lb,
		e: z.i - z.lb,

		v: string(z.Source[z.p:z.i]),
		t: t,
	})

	z.p = z.i

	switch t {
	case TokenEol:
		z.l++
		z.li = 0
		z.lb = z.i
	}
}

func isTerminating(c rune) bool {
	switch c {
	case '\r':
	case '\n':
	case ' ':
	case '\t':
	case ';':
	default:
		return false
	}

	return true
}

type lexerError struct {
	l  int
	li int

	msg string
}

func (e lexerError) Line() int {
	return e.l
}

func (e lexerError) Begin() int {
	return e.li
}

func (e lexerError) End() int {
	return e.li
}

func (e lexerError) String() string {
	return e.msg
}

func (z *Lexer) readValue() {
	p, n := utf8.DecodeRune(z.Source[z.i:])
	if p == '"' {
		z.inc(n)
		for {
			if z.i == len(z.Source) {
				z.fail("incomplete string")
				return
			}

			c, n := utf8.DecodeRune(z.Source[z.i:])
			if c == '"' && p != '\\' {
				z.inc(n)

				s := string(z.Source[z.p+1 : z.i-1])
				v := "\"" + strings.ReplaceAll(s, "\\\"", "\"") + "\""

				z.ts = append(z.ts, Token{
					l: z.l,
					b: z.p - z.lb,
					e: z.i - z.lb,

					v: v,
					t: TokenValue,
				})

				z.p = z.i
				return
			}

			z.inc(n)

			p = c
		}
	} else {
		for {
			if z.i == len(z.Source) {
				z.emit(TokenValue)
				return
			}

			c, n := utf8.DecodeRune(z.Source[z.i:])
			if isTerminating(c) {
				z.emit(TokenValue)
				return
			}

			z.inc(n)
		}
	}
}

func (z *Lexer) skipWhitespace() {
	for {
		if z.i == len(z.Source) {
			return
		}

		c, n := utf8.DecodeRune(z.Source[z.i:])
		if c != ' ' && c != '\t' {
			return
		}

		z.inc(n)
		z.p = z.i
	}
}

func (z *Lexer) readComment() {
	for {
		l := z.i - z.p
		if z.i == len(z.Source) {
			if l < 2 {
				z.fail("incomplete comment")
				return
			}

			z.ts = append(z.ts, Token{
				l: z.l,
				b: z.p - z.lb,
				e: z.i - z.lb,

				v: string(z.Source[z.p+2 : z.i]),
				t: TokenComment,
			})

			z.p = z.i
			return
		}

		c, n := utf8.DecodeRune(z.Source[z.i:])
		if l < 2 {
			if c != '/' {
				z.fail("incomplete comment")
				return
			}
		} else {
			if c == '\r' || c == '\n' {
				z.ts = append(z.ts, Token{
					l: z.l,
					b: z.p - z.lb,
					e: z.i - z.lb,

					v: string(z.Source[z.p+2 : z.i]),
					t: TokenComment,
				})

				z.p = z.i
				return
			}
		}

		z.inc(n)
	}
}

func (z *Lexer) readSeparator() {
	for {
		if z.i == len(z.Source) {
			z.emit(TokenEol)
			return
		}

		c, n := utf8.DecodeRune(z.Source[z.i:])
		switch c {
		case '\r':
			z.inc(n)
			if z.i < len(z.Source) {
				c2, n2 := utf8.DecodeRune(z.Source[z.i:])
				if c2 == '\n' {
					z.inc(n2)
				}
			}
			z.emit(TokenEol)
			return
		case '\n':
			z.inc(n)
			z.emit(TokenEol)
			return
		case ';':
			z.inc(n)
			z.emit(TokenSemicolon)
			return
		}

		z.inc(n)
	}
}

func (z *Lexer) readTokens() {
	z.skipWhitespace()

	if z.i < len(z.Source) {
		c, n := utf8.DecodeRune(z.Source[z.i:])

		var nc rune
		if z.i+n < len(z.Source) {
			nc, _ = utf8.DecodeRune(z.Source[z.i+n:])
		}

		if c == '/' && nc == '/' {
			z.readComment()
		} else if c == '\n' || c == '\r' || c == ';' {
			z.readSeparator()
		} else {
			z.readValue()
		}
	}
}

func (z *Lexer) read() {
	if z.i == len(z.Source) {
		return
	}

	z.readTokens()
}

func (z *Lexer) Return() bool {
	if z.tsi == 0 {
		return false
	}

	z.tsi--

	return true
}

func (z *Lexer) Scan() bool {
	if len(z.ts) > z.tsi {
		z.tsi++
	}

	if len(z.ts) == z.tsi {
		z.read()
	}

	return len(z.ts) > z.tsi
}

func (z *Lexer) Curr() Token {
	if len(z.ts) == z.tsi {
		return Token{}
	}

	return z.ts[z.tsi]
}

func (z *Lexer) Errors() []lexerError {
	return z.diag
}
