package teal

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/dragmz/tealex"
	"github.com/pkg/errors"
)

type ParserContext interface {
	Read_int8() int8
	Read_uint8() uint8
	Read_int16() int16
	Read_rbyte() []byte
	Read_rrbyte() [][]byte
	Read() tealex.Token
}

func (a *parserContext) Read_int8() int8 {
	return a.mustReadInt8()
}

func (a *parserContext) Read_uint8() uint8 {
	return a.mustReadUint8()
}

func (a *parserContext) Read_int16() int16 {
	return a.mustReadInt16()
}

func (a *parserContext) Read_rbyte() []byte {
	return a.mustReadBytes()
}

func (a *parserContext) Read_rrbyte() [][]byte {
	return a.readBytesArray()
}

func (a *parserContext) Read() tealex.Token {
	a.mustRead()
	return a.args.Curr()
}

type parseError struct {
	error

	l int
	b int
	e int
}

func (e parseError) Line() int {
	return e.l
}

func (e parseError) Begin() int {
	return e.b
}

func (e parseError) End() int {
	return e.e
}

func (e parseError) String() string {
	return e.error.Error()
}

func readInt8(s string) (int8, error) {
	v, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0, err
	}

	return int8(v), nil
}

func readUint8(s string) (uint8, error) {
	v, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(v), nil
}

func readInt16(s string) (int16, error) {
	v, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, err
	}

	return int16(v), nil
}

func readInt(a *arguments) (uint64, error) {
	val, err := strconv.ParseUint(a.Text(), 0, 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func (c *parserContext) mustReadArg() {
	if !c.args.Scan() {
		c.failPrev(errors.Errorf("missing arg"))
	}
}

func (c *parserContext) parseUint64() uint64 {
	v, err := readInt(c.args)
	if err != nil {
		c.failCurr(errors.Wrapf(err, "failed to parse uint64"))
	}

	return v
}

func (c *parserContext) parseBytes() []byte {
	arg := c.args.Curr().String()

	if strings.HasPrefix(arg, "base32(") || strings.HasPrefix(arg, "b32(") {
		close := strings.IndexRune(arg, ')')
		if close == -1 {
			c.failCurr(errors.New("byte base32 arg lacks close paren"))
		}

		open := strings.IndexRune(arg, '(')
		val, err := base32DecodeAnyPadding(arg[open+1 : close])
		if err != nil {
			c.failCurr(err)
		}

		return val
	}

	if strings.HasPrefix(arg, "base64(") || strings.HasPrefix(arg, "b64(") {
		close := strings.IndexRune(arg, ')')
		if close == -1 {
			c.failCurr(errors.New("byte base64 arg lacks close paren"))
		}

		open := strings.IndexRune(arg, '(')
		val, err := base64.StdEncoding.DecodeString(arg[open+1 : close])
		if err != nil {
			c.failCurr(err)
		}
		return val
	}

	if strings.HasPrefix(arg, "0x") {
		val, err := hex.DecodeString(arg[2:])
		if err != nil {
			c.failCurr(err)
		}
		return val
	}

	if arg == "base32" || arg == "b32" {
		l := c.mustRead()
		val, err := base32DecodeAnyPadding(l)
		if err != nil {
			c.failCurr(err)
		}

		return val
	}

	if arg == "base64" || arg == "b64" {
		l := c.mustRead()
		val, err := base64.StdEncoding.DecodeString(l)
		if err != nil {
			c.failCurr(err)
		}

		return val
	}

	if len(arg) > 1 && arg[0] == '"' && arg[len(arg)-1] == '"' {
		val, err := parseStringLiteral(arg)
		if err != nil {
			c.failCurr(err)
		}
		return val
	}

	c.failCurr(fmt.Errorf("byte arg did not parse: %v", arg))

	return nil
}

func (c *parserContext) parseUint8() uint8 {
	v, err := readUint8(c.args.Text())
	if err != nil {
		c.failCurr(errors.Wrapf(err, "failed to parse uint8"))
	}

	return v
}

func (c *parserContext) parseInt16() int16 {
	v, err := readInt16(c.args.Text())
	if err != nil {
		c.failCurr(errors.Wrapf(err, "failed to parse uint8"))
	}

	return v
}

func (c *parserContext) readBytesArray() [][]byte {
	res := [][]byte{}

	for c.args.Scan() {
		bs := c.parseBytes()
		res = append(res, bs)
	}

	return res
}

func (c *parserContext) mustReadBytes() []byte {
	c.mustReadArg()
	return c.parseBytes()
}

func (c *parserContext) mustReadInt() uint64 {
	c.mustReadArg()
	return c.parseUint64()
}

func (c *parserContext) mustReadUint8() uint8 {
	c.mustReadArg()
	return c.parseUint8()
}

func (c *parserContext) mustReadInt16() int16 {
	c.mustReadArg()
	return c.parseInt16()
}

func (c *parserContext) maybeReadArg() bool {
	return c.args.Scan()
}

func (c *parserContext) maybeReadUint8() (uint8, bool) {
	ok := c.maybeReadArg()
	if !ok {
		return 0, false
	}

	return c.parseUint8(), true
}

func (c *parserContext) parseInt8() int8 {
	v, err := readInt8(c.args.Text())
	if err != nil {
		c.failCurr(errors.Wrapf(err, "failed to parse int8"))
	}

	return v
}

func (c *parserContext) mustReadInt8() int8 {
	c.mustReadArg()
	return c.parseInt8()
}

func (c *parserContext) mustRead() string {
	c.mustReadArg()
	return c.args.Text()
}

type parserContext struct {
	args *arguments
}

type recoverable struct{}

func (c *parserContext) failAt(l int, b int, e int, err error) {
	c.errorAt(l, b, e, err)
	panic(recoverable{})
}

func (c *parserContext) errorAt(l int, b int, e int, err error) {
	panic(parseError{l: l, b: b, e: e, error: err})
}

func (c *parserContext) failToken(t tealex.Token, err error) {
	c.failAt(t.Line(), t.Begin(), t.End(), err)
}

func (c *parserContext) failCurr(err error) {
	c.failToken(c.args.Curr(), err)
}

func (c *parserContext) failPrev(err error) {
	c.failToken(c.args.Prev(), err)
}

type Subline struct {
	Tokens []tealex.Token
}

type Line struct {
	Tokens []tealex.Token
	Subs   []Subline
}

func readLines(ts []tealex.Token) []Line {
	lines := []Line{}

	p := 0
	for i := 0; i < len(ts); i++ {
		t := ts[i]

		j := i + 1
		eol := t.Type() == tealex.TokenEol

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
			case tealex.TokenSemicolon:
				l.Subs = append(l.Subs, Subline{
					Tokens: l.Tokens[sf:i],
				})
				sf = i + 1
			case tealex.TokenComment:
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

	lexer := &tealex.Lexer{
		Source: []byte(src),
	}

	var tokens []tealex.Token

	for lexer.Scan() {
		tokens = append(tokens, lexer.Curr())
	}

	lines := readLines(tokens)
	prepareLines(lines)

	for _, line := range lines {
		for _, sub := range line.Subs {
			var op TealOp

			ctx := &parserContext{args: &arguments{ts: sub.Tokens}}

			f := ctx.Read()
			v := f.String()

			if strings.HasSuffix(v, "#pragma") {
				name := ctx.Read()

				if name.String() != "version" {
					panic(fmt.Sprintf("unsupported #pragma: '%s'", name.String()))
				}

				version := ctx.Read_uint8()

				op = &Teal_pragma_version{
					Version: int(version),
				}
			} else if strings.HasSuffix(v, ":") {
				op = &Teal_label{Name: v[:len(v)-1]}
			} else {
				switch v {
				case "int":
					op = Parse_Teal_int_op(ctx)
				case "byte":
					op = Parse_Teal_byte_op(ctx)
				default:
					pa := &parserContext{&arguments{ts: sub.Tokens}}
					op = parseTeal(pa)
				}
			}
			t = append(t, op)
		}
	}

	return t
}
