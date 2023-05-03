package compiler

import (
	"ceal/parser"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/pkg/errors"
)

type cealPreprocessor struct {
	includes  *cealIncludes
	processed map[string]bool
	defines   map[string]string

	stack *defStack
}

type defStack struct {
	parent *defStack
	skip   bool
}

func (d *defStack) push() *defStack {
	return &defStack{
		skip:   d.skip,
		parent: d,
	}
}

func (d *defStack) pop() *defStack {
	if d == nil {
		panic("#endif without #if")
	}
	return d.parent
}

func (p *cealPreprocessor) preprocess(name string, src string) (string, error) {
	src = strings.ReplaceAll(strings.ReplaceAll(src, "\r\n", "\n"), "\r", "\n")

	p.stack = p.stack.push()
	defer func() {
		p.stack = p.stack.pop()
	}()

	lines := strings.Split(src, "\n")

	i := 0

	for i < len(lines) {
		clear := func() {
			lines[i] = ""
		}

		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			clear()
			i += 1
			continue
		}

		parts := strings.SplitN(strings.TrimSpace(line), " ", 2)

		switch parts[0] {
		case "#endif":
			p.stack = p.stack.pop()
			clear()
		case "#ifndef":
			if len(parts) < 2 {
				return "", fmt.Errorf("invalid #ifndef: '%s'", line)
			}

			p.stack = p.stack.push()

			l := parser.NewCLexer(antlr.NewInputStream(parts[1]))
			l.RemoveErrorListeners()
			t := l.NextToken()

			if t.GetTokenType() != parser.CLexerID {
				return "", fmt.Errorf("invalid #ifndef: '%s'", line)
			}

			if _, ok := p.defines[t.GetText()]; ok {
				p.stack.skip = true
			}

			clear()
		case "#ifdef":
			if len(parts) < 2 {
				return "", fmt.Errorf("invalid #ifndef: '%s'", line)
			}

			p.stack = p.stack.push()

			l := parser.NewCLexer(antlr.NewInputStream(parts[1]))
			l.RemoveErrorListeners()
			t := l.NextToken()

			if t.GetTokenType() != parser.CLexerID {
				return "", fmt.Errorf("invalid #ifdef: '%s'", line)
			}

			if _, ok := p.defines[t.GetText()]; !ok {
				p.stack.skip = true
			}

			clear()
		default:
			if p.stack.skip {
				clear()
			} else {
				stop := false

				for !stop {
					stop = true
					l := parser.NewCLexer(antlr.NewInputStream(line))
					l.RemoveErrorListeners()

				inner:

					for {
						t := l.NextToken()
						if t.GetTokenType() == antlr.TokenEOF {
							break
						}

						if t.GetTokenType() == parser.CLexerID {
							id := t.GetText()
							if v, ok := p.defines[id]; ok && len(v) > 0 {
								// replace token from start to stop with define value
								line = line[:t.GetStart()] + v + line[t.GetStop()+1:]
								lines[i] = line
								stop = false
								break inner
							}
						}
					}
				}

				parts := strings.SplitN(strings.TrimSpace(line), " ", 2)

				switch parts[0] {
				case "#if":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #if: '%s'", line)
					}

					p.stack = p.stack.push()

					l := parser.NewCLexer(antlr.NewInputStream(parts[1]))
					l.RemoveErrorListeners()
					t := l.NextToken()

					if t.GetTokenType() != parser.CLexerINT {
						return "", fmt.Errorf("invalid #if: '%s'", line)
					}

					if t.GetText() == "0" {
						p.stack.skip = true
					}

					clear()

				case "#include":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #include: '%s'", line)
					}

					quoted := strings.TrimSpace(parts[1])

					inc, ipath, err := p.includes.resolve(quoted)
					if err != nil {
						return "", errors.Wrapf(err, "failed to resolve #include file: '%s'", name)
					}

					inc, err = p.preprocess(ipath, inc)
					if err != nil {
						return "", errors.Wrapf(err, "failed to preprocess #include file: '%s'", name)
					}

					inclines := strings.Split(inc, "\n")

					// insert inclines at current line removing the current line
					lines = append(lines[:i], append(inclines, lines[i+1:]...)...)
				case "#define":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #define: '%s'", line)
					}

					l := parser.NewCLexer(antlr.NewInputStream(parts[1]))
					l.RemoveErrorListeners()
					t := l.NextToken()

					if t.GetTokenType() != parser.CLexerID {
						return "", fmt.Errorf("invalid #define: '%s'", line)
					}

					id := t.GetText()

					items := strings.SplitN(parts[1], " ", 2)

					switch len(items) {
					case 1:
						p.defines[id] = ""
					case 2:
						value := parts[1][t.GetStop()+1:]
						p.defines[id] = strings.TrimSpace(value)
					}

					clear()
				case "#undef":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #undef: '%s'", line)
					}

					name := strings.TrimSpace(parts[1])

					delete(p.defines, name)

					clear()
				case "#pragma":
					if len(parts) < 2 {
						return "", fmt.Errorf("invalid #pragma: '%s'", line)
					}

					switch parts[1] {
					case "once":
						clear()
						if p.processed[name] {
							p.stack.skip = true
						} else {
							p.processed[name] = true
						}
					}
				}
			}
		}

		i++
	}

	src = strings.Join(lines, "\n")

	return src, nil
}
