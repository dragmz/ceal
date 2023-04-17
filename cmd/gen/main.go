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
package compiler

const AvmMainName = "avm_main"
`)

	fmt.Fprintf(bw, "const AvmVersion = %d\n", cs.Version)

	bw.WriteString(`
type AstStatement interface {
	String() string
}

type BuiltinFunctionParamData struct
{
	t string
	name string;
};

type BuiltinFunctionData struct
{
	t string
	name string
	op string

	stack []BuiltinFunctionParamData
	imm []BuiltinFunctionParamData
	returns int
};
`)

	bw.WriteString(`
var builtin_functions = []BuiltinFunctionData {
`)

	for _, op := range cs.Ops {
		name := ceal.FormatOpName(op.Name)
		ret := ceal.ReadReturnTypeName(op)

		bw.WriteString("\t{\n")
		fmt.Fprintf(bw, "\t\tt: \"%s\", name: \"avm_%s\", op: \"%s\",\n", ret, name, op.Name)
		bw.WriteString("\t\tstack: []BuiltinFunctionParamData{\n")
		for _, arg := range op.Stacks {
			fmt.Fprintf(bw, "\t\t\t{ t: \"%s\", name: \"%s\" },\n", arg.Type, arg.Name)
		}
		bw.WriteString("\t\t},\n")
		bw.WriteString("\t\timm: []BuiltinFunctionParamData{\n")
		for _, arg := range op.Imms {
			fmt.Fprintf(bw, "\t\t\t{ t: \"%s\", name: \"%s\" },\n", arg.Type, arg.Name)
		}
		bw.WriteString("\t\t},\n")
		fmt.Fprintf(bw, "\t\treturns: %d,\n", len(op.Returns))
		bw.WriteString("\t},\n")
		//
		bw.WriteString("\t{\n")
		fmt.Fprintf(bw, "\t\tt: \"%s\", name: \"avm_%s_op\", op: \"%s\",\n", ret, name, op.Name)
		bw.WriteString("\t\timm: []BuiltinFunctionParamData{\n")
		for _, arg := range op.Imms {
			fmt.Fprintf(bw, "\t\t\t{ t: \"%s\", name: \"%s\" },\n", arg.Type, arg.Name)
		}
		bw.WriteString("\t\t},\n")
		bw.WriteString("\t},\n")
	}

	bw.WriteString(`};
`)

	bw.WriteString(`
type BuiltinStructFieldData struct
{
	t string
	name string
	fun string
};

type BuiltinStructFunctionParamData struct
{
	t string
	name string
};

type BuiltinStructFunctionData struct
{
	t string
	name string
	fun string
	params []BuiltinStructFunctionParamData
};
	
type BuiltinStructData struct
{
	name string
	fields []BuiltinStructFieldData
	functions []BuiltinStructFunctionData
};
	
var builtin_structs = []BuiltinStructData {
`)

	for _, op := range cs.Ops {
		name := ceal.FormatOpName(op.Name)

		switch len(op.Imms) {
		case 1:
			if len(op.Enum) == 0 {
				break
			}
			bw.WriteString("\t{\n")
			fmt.Fprintf(bw, "\t\tname: \"avm_%s_t\",\n", name)
			bw.WriteString("\t\tfields: []BuiltinStructFieldData{\n")
			for _, e := range op.Enum {
				fmt.Fprintf(bw, "\t\t\t{t: \"%s\", name: \"%s\", fun: \"avm_%s\"},\n", e.Type, e.Name, name)
			}
			bw.WriteString("\t\t},\n")
			bw.WriteString("\t\tfunctions: []BuiltinStructFunctionData{\n")
			bw.WriteString("\t\t},\n")
			bw.WriteString("\t},\n")
		default:
			if len(op.Enum) == 0 {
				break
			}
			name := ceal.FormatOpName(op.Name)

			bw.WriteString("\t{\n")
			fmt.Fprintf(bw, "\t\tname: \"avm_%s_t\",\n", name)
			bw.WriteString("\t\tfields: []BuiltinStructFieldData{\n")
			bw.WriteString("\t\t},\n")
			bw.WriteString("\t\tfunctions: []BuiltinStructFunctionData{\n")
			for _, e := range op.Enum {
				bw.WriteString("\t\t\t{\n")
				fmt.Fprintf(bw, "\t\t\t\tt: \"%s\", name: \"%s\", fun: \"avm_%s\",\n", e.Type, e.Name, name)
				bw.WriteString("\t\t\t\tparams: []BuiltinStructFunctionParamData{\n")
				for i := 1; i < len(op.Imms); i++ {
					imm := op.Imms[i]
					fmt.Fprintf(bw, "\t\t\t\t\t{ t: \"%s\", name: \"%s\" },\n", imm.Type, imm.Name)
				}
				bw.WriteString("\t\t\t\t},\n")
				bw.WriteString("\t\t\t},\n")
			}
			bw.WriteString("\t\t},\n")
			bw.WriteString("\t},\n")
		}

		switch len(op.Returns) {
		case 0:
		case 1:
		default:
			name := ceal.FormatOpName(op.Name)
			rt := ceal.ReadReturnTypeName(op)
			bw.WriteString("\t{\n")
			fmt.Fprintf(bw, "\t\tname: \"%s\",\n", rt)
			bw.WriteString("\t\tfields: []BuiltinStructFieldData{\n")
			for _, r := range op.Returns {
				fmt.Fprintf(bw, "\t\t\t{t: \"%s\", name: \"%s\", fun: \"avm_%s\",},\n", r.Type, r.Name, name)
			}
			bw.WriteString("\t\t},\n")
			bw.WriteString("\t\tfunctions: []BuiltinStructFunctionData{\n")
			bw.WriteString("\t\t},\n")
			bw.WriteString("\t},\n")
		}
	}

	bw.WriteString(`
};
`)

	bw.WriteString(`
type BuiltinVariableData struct
{
	t string
	name string
};
`)

	bw.WriteString(`
var builtin_variables = []BuiltinVariableData {
`)

	for _, op := range cs.Ops {
		switch len(op.Imms) {
		case 0:
		default:
			if len(op.Enum) == 0 {
				break
			}
			name := ceal.FormatOpName(op.Name)

			fmt.Fprintf(bw, "\t{ \"avm_%s_t\", \"avm_%s\" },\n", name, op.Name)
		}
	}
	fmt.Fprintf(bw, "\t%s,\n", `{ "uint64", "NoOp" }`)
	fmt.Fprintf(bw, "\t%s,\n", `{ "uint64", "OptIn" }`)
	fmt.Fprintf(bw, "\t%s,\n", `{ "uint64", "DeleteApplication" }`)
	fmt.Fprintf(bw, "\t%s,\n", `{ "uint64", "UpdateApplication" }`)
	fmt.Fprintf(bw, "\t%s,\n", `{ "uint64", "CloseOut" }`)
	fmt.Fprintf(bw, "\t%s,\n", `{ "uint64", "ClearState" }`)

	bw.WriteString(`
};
`)

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
