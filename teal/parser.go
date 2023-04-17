package teal

import (
	"fmt"
	"strings"
)

type ParserToken interface {
	Value() string
}

type parserToken struct {
	t Token
}

func (t *parserToken) Value() string {
	return t.t.String()
}

type ParserArgs interface {
	Read_int8() int8
	Read_uint8() uint8
	Read_int16() int16
	Read_rbyte() []byte
	Read_rrbyte() [][]byte
	Read() *Token
}

type parserArgs struct {
	items []Token
	index int
}

func (a *parserArgs) Read_int8() int8 {
	panic("not implemented")
}

func (a *parserArgs) Read_uint8() uint8 {
	panic("not implemented")
}

func (a *parserArgs) Read_int16() int16 {
	panic("not implemented")
}

func (a *parserArgs) Read_rbyte() []byte {
	panic("not implemented")
}

func (a *parserArgs) Read_rrbyte() [][]byte {
	panic("not implemented")
}

func (a *parserArgs) Read() *Token {
	if a.index >= len(a.items) {
		return nil
	}

	t := a.items[a.index]
	a.index++

	return &t
}

type Subline struct {
	Tokens []Token
}

type Line struct {
	Tokens []Token
	Subs   []Subline
}

func readLines(ts []Token) []Line {
	lines := []Line{}

	p := 0
	for i := 0; i < len(ts); i++ {
		t := ts[i]

		j := i + 1
		eol := t.Type() == TokenEol

		if eol || j == len(ts) {
			k := j
			if eol {
				k--
			}

			lines = append(lines, Line{
				Tokens: ts[p:k],
			})
			p = j
		}
	}

	return lines
}

func prepareLines(lines []Line) {
	for li, l := range lines {
		sf := 0

		for i := 0; i < len(l.Tokens); i++ {
			t := l.Tokens[i]
			switch t.Type() {
			case TokenSemicolon:
				l.Subs = append(l.Subs, Subline{
					Tokens: l.Tokens[sf:i],
				})
				sf = i + 1
			case TokenComment:
				l.Tokens = l.Tokens[:i]
				l.Subs = append(l.Subs, Subline{
					Tokens: l.Tokens[sf:i],
				})
				sf = i
			}
		}

		if sf <= len(l.Tokens) {
			l.Subs = append(l.Subs, Subline{
				Tokens: l.Tokens[sf:len(l.Tokens)],
			})
		}
		lines[li] = l
	}
}

func ParseTeal(src string) Teal {
	var t Teal

	lexer := &Lexer{
		Source: []byte(src),
	}

	var tokens []Token

	for lexer.Scan() {
		tokens = append(tokens, lexer.Curr())
	}

	lines := readLines(tokens)
	prepareLines(lines)

	for _, line := range lines {
		for _, sub := range line.Subs {
			var op TealOp

			fpa := &parserArgs{items: sub.Tokens}
			f := fpa.Read()
			v := f.String()

			if strings.HasSuffix(v, "#pragma") {
				name := fpa.Read()

				if name.String() != "version" {
					panic(fmt.Sprintf("unsupported #pragma: '%s'", name.String()))
				}

				version := fpa.Read_uint8()

				op = &Teal_pragma_version{
					Version: int(version),
				}
			} else if strings.HasSuffix(v, ":") {
				op = &Teal_label{Name: v[:len(v)-1]}
			} else {
				switch v {
				case "int":
					op = Parse_Teal_int_op(fpa)
				default:
					pa := &parserArgs{items: sub.Tokens}
					op = parseTeal(pa)
				}
			}
			t = append(t, op)
		}
	}

	return t
}
