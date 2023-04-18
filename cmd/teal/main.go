package main

import (
	"bufio"
	"bytes"
	"ceal"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type args struct {
	Spec string
	Out  string
}

func readGoType(arg ceal.CealArg) string {
	var t string
	switch arg.Type {
	case "uint8":
		t = arg.Type
	case "int8":
		t = arg.Type
	case "int16":
		t = arg.Type
	case "uint16":
		t = arg.Type
	case "uint64":
		t = arg.Type
	case "bytes":
		t = "[]byte"
	default:
		panic(fmt.Sprintf("unsupported arg type: '%s'", arg.Type))
	}

	if arg.Array {
		t = fmt.Sprintf("[]%s", t)
	}

	return t
}

func makeAvmCpp(cs ceal.CealSpec, bw *bufio.Writer) error {
	defer bw.Flush()

	/*
		struct BuiltinFunction
		{
			std::string type;
			std::string name;
			std::string op;

			std::size_t stack;
			std::size_t imm;
		};

		std::vector<BuiltinFunction> builtin_functions
		{
			{ "uint64", "avm_box_create", "box_create", 2, 0},
			{ "void", "avm_box_put", "box_put", 2, 0 },
			{ "bytes", "avm_itob", "itob", 1, 0 },
			{ "bytes", "avm_box_extract", "box_extract", 3, 0 },
			{ "void", "avm_log", "log", 1, 0 },
			{ "void", "avm_err", "err", 0, 0 },
		};
	*/

	bw.WriteString(`
package teal

import (
	"strings"
	"fmt"
)
`)

	bw.WriteString(`
type TealAst interface
{
	Teal() Teal
};

`)

	for _, op := range cs.Ops {
		name := ceal.FormatOpName(op.Name)
		fmt.Fprintf(bw, "type Teal_%s struct {\n", name)
		for _, arg := range op.Stacks {
			fmt.Fprintf(bw, "\t%s TealAst\n", strings.ToUpper(arg.Name))
		}
		for _, arg := range op.Imms {
			fmt.Fprintf(bw, "\t%s %s\n", arg.Name, readGoType(arg))
		}

		bw.WriteString("}\n")

		fmt.Fprintf(bw, "func (a *Teal_%s) Teal() Teal {\n", name)
		bw.WriteString("\treturn Teal{a}\n")
		bw.WriteString("}\n")

		fmt.Fprintf(bw, "func (a *Teal_%s) String() string {\n", name)
		bw.WriteString("\tres := strings.Builder{}\n")

		for _, arg := range op.Stacks {
			fmt.Fprintf(bw, "\tfor _, op := range a.%s.Teal() {\n", strings.ToUpper(arg.Name))
			bw.WriteString("\t\tres.WriteString(op.String())\n")
			bw.WriteString("\t\tres.WriteString(\"\\n\")\n")
			bw.WriteString("\t}\n")
		}

		fmt.Fprintf(bw, "\tres.WriteString(\"%s\")\n", op.Name)
		if len(op.Imms) > 0 {
			for _, arg := range op.Imms {
				bw.WriteString("\tres.WriteString(\" \")\n")

				// TODO: handle other types than just ints
				fmt.Fprintf(bw, "\tres.WriteString(fmt.Sprintf(\"%%d\", a.%s))\n", arg.Name)
			}
		}

		bw.WriteString("\treturn res.String()\n")
		bw.WriteString("}\n")

		fmt.Fprintf(bw, "type Teal_%s_op struct {\n", name)
		for _, arg := range op.Imms {
			fmt.Fprintf(bw, "\t%s %s\n", arg.Name, readGoType(arg))
		}
		bw.WriteString("}\n")

		fmt.Fprintf(bw, "func Parse_Teal_%s_op(ctx ParserContext) *Teal_%s_op {\n", name, name)
		fmt.Fprintf(bw, "\tt := &Teal_%s_op{\n", name)
		for _, imm := range op.Imms {
			fmt.Fprintf(bw, "\t\t%s: ctx.Read_%s(),\n", imm.Name, strings.ReplaceAll(readGoType(imm), "[]", "r"))
		}
		bw.WriteString("\t}\n")
		bw.WriteString("\treturn t\n")
		bw.WriteString("}\n")

		bw.WriteString("\n")

		fmt.Fprintf(bw, "func (a *Teal_%s_op) String() string {\n", name)
		bw.WriteString("\tres := strings.Builder{}\n")

		fmt.Fprintf(bw, "\tres.WriteString(\"%s\")\n", op.Name)
		if len(op.Imms) > 0 {
			for _, arg := range op.Imms {
				bw.WriteString("\tres.WriteString(\" \")\n")

				// TODO: handle other types than just ints
				fmt.Fprintf(bw, "\tres.WriteString(fmt.Sprintf(\"%%d\", a.%s))\n", arg.Name)
			}
		}
		bw.WriteString("\treturn res.String()")
		bw.WriteString("}\n")
	}

	bw.WriteString("\n")

	bw.WriteString("func parseTeal(ctx ParserContext) TealOp {\n")
	bw.WriteString("\tfirst := ctx.Read()\n")
	bw.WriteString("\tvalue := first.String()\n")
	bw.WriteString("\tswitch value {\n")
	for _, op := range cs.Ops {
		name := ceal.FormatOpName(op.Name)
		fmt.Fprintf(bw, "\t\tcase \"%s\":\n", op.Name)
		fmt.Fprintf(bw, "\t\treturn Parse_Teal_%s_op(ctx)\n", name)
	}
	bw.WriteString("\t\tdefault:\n")
	bw.WriteString("\t\t\tpanic(fmt.Sprintf(\"unexpected op: %s\", value))")
	bw.WriteString("\t}\n")
	bw.WriteString("}\n")

	return nil
}

func run(a args) error {
	sf, err := os.Open(a.Spec)
	if err != nil {
		return errors.Wrap(err, "failed to open spec file")
	}

	defer sf.Close()

	var cs ceal.CealSpec
	d := json.NewDecoder(sf)
	err = d.Decode(&cs)
	if err != nil {
		return errors.Wrap(err, "failed to decode spec")
	}

	buf := bytes.Buffer{}

	bw := bufio.NewWriter(&buf)
	err = makeAvmCpp(cs, bw)
	if err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "failed to format source")
	}

	f, err := os.Create(a.Out)
	if err != nil {
		return errors.Wrap(err, "failed to create file")
	}

	defer f.Close()

	f.Write(formatted)

	return nil
}

func main() {
	var a args
	flag.StringVar(&a.Spec, "spec", "", "spec file path")
	flag.StringVar(&a.Out, "out", "", "output file path")
	flag.Parse()

	err := run(a)
	if err != nil {
		panic(err)
	}
}
