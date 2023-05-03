package compiler

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

type cealPreprocessor struct {
	includes  []string
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
	if d.parent == nil {
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

		parts := strings.SplitN(strings.TrimSpace(line), " ", 2)

		if len(parts) == 0 {
			clear()
			i += 1
			continue
		}

		switch parts[0] {
		case "#endif":
			p.stack = p.stack.pop()
			clear()
		case "#ifndef":
			if len(parts) < 2 {
				return "", fmt.Errorf("invalid #ifndef: '%s'", line)
			}

			p.stack = p.stack.push()

			name := strings.TrimSpace(parts[1])

			if _, ok := p.defines[name]; ok {
				p.stack.skip = true
			}

			clear()
		case "#ifdef":
			if len(parts) < 2 {
				return "", fmt.Errorf("invalid #ifdef: '%s'", line)
			}

			p.stack = p.stack.push()

			name := strings.TrimSpace(parts[1])

			if _, ok := p.defines[name]; !ok {
				p.stack.skip = true
			}

			clear()
		default:
			if p.stack.skip {
				clear()
			} else {
				switch parts[0] {
				case "#include":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #include: '%s'", line)
					}

					quoted := strings.TrimSpace(parts[1])

					begin := string(quoted[0])
					var end string

					switch begin {
					case "\"":
						end = begin
					case "<":
						end = ">"
					default:
						return "", fmt.Errorf("unexpected #include quote in: '%s'", line)
					}

					if end != string(quoted[len(quoted)-1]) {
						return "", fmt.Errorf("#include quote mismatch: '%s'", line)
					}

					name := quoted[1 : len(quoted)-1]

					inc, ipath, err := func() (string, string, error) {
						for _, dir := range p.includes {
							ip := path.Join(dir, name)

							bs, err := os.ReadFile(ip)
							if os.IsNotExist(err) {
								continue
							}

							if err != nil {
								return "", "", errors.Wrapf(err, "failed to read #include file: '%s'", name)
							}

							inc := string(bs)
							return inc, ip, nil
						}

						return "", "", errors.Errorf("failed to find #include file: '%s'", name)
					}()

					if err != nil {
						return "", err
					}

					inc, err = p.preprocess(ipath, inc)
					if err != nil {
						return "", err
					}

					inclines := strings.Split(inc, "\n")

					// insert inclines at current line removing the current line
					lines = append(lines[:i], append(inclines, lines[i+1:]...)...)
				case "#define":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #define: '%s'", line)
					}

					items := strings.SplitN(parts[1], " ", 2)

					switch len(items) {
					case 1:
						p.defines[strings.TrimSpace(items[0])] = ""
					case 2:
						p.defines[strings.TrimSpace(items[0])] = strings.TrimSpace(items[1])
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
