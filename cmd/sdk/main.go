package main

import (
	"bufio"
	"ceal"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type args struct {
	Spec string
	Out  string
}

func formatParams(op ceal.CealSpecOp, skipImmsCount int) string {
	b := strings.Builder{}

	for _, arg := range op.Stacks {
		if b.Len() > 0 {
			b.WriteString(", ")
		}

		b.WriteString("STACK ")

		b.WriteString(arg.Type)
		b.WriteString(" ")
		b.WriteString(arg.Name)
	}

	for i := skipImmsCount; i < len(op.Imms); i++ {
		arg := op.Imms[i]
		if b.Len() > 0 {
			b.WriteString(", ")
		}

		b.WriteString("IMMEDIATE ")

		b.WriteString(arg.Type)
		b.WriteString(" ")
		b.WriteString(arg.Name)
	}

	return b.String()
}

func getType(r byte) string {
	switch r {
	case '.':
		return "any"
	case 'U':
		return "uint64"
	case 'B':
		return "bytes"
	default:
		panic(fmt.Sprintf("unhandled: %s", string(r)))
	}

}

func run(a args) error {
	r, err := os.Open(a.Spec)
	if err != nil {
		return errors.Wrap(err, "failed to download langspec")
	}

	defer r.Close()

	var cs ceal.CealSpec

	d := json.NewDecoder(r)
	err = d.Decode(&cs)
	if err != nil {
		return errors.Wrap(err, "failed to decode langspec")
	}

	w, err := os.Create(a.Out)
	if err != nil {
		return errors.Wrap(err, "failed to create lanspec output file")
	}

	defer w.Close()

	bw := bufio.NewWriter(w)

	defer bw.Flush()

	bw.WriteString(`
#pragma once

#include <variant>

#define IMMEDIATE
#define STACK

using uint64 = unsigned long long;
using bytes = std::variant<const char*, const unsigned char*>;

const uint64 NoOp = 0;
const uint64 OptIn = 1;
const uint64 CloseOut = 2;
const uint64 ClearState = 3;
const uint64 UpdateApplication = 4;
const uint64 DeleteApplication = 5;

struct any
{
	bool operator==(const any);
};

`)

	for _, op := range cs.Ops {
		name := ceal.FormatOpName(op.Name)

		rt := ceal.ReadReturnTypeName(op)

		switch len(op.Returns) {
		case 0:
		case 1:
		default:
			bw.WriteString(fmt.Sprintf("struct %s\n", rt))
			bw.WriteString("{\n")
			for _, r := range op.Returns {
				bw.WriteString(fmt.Sprintf("\t%s %s;\n", r.Type, r.Name))
			}
			bw.WriteString("};\n")
		}

		bw.WriteString("/*\n")
		bw.WriteString(fmt.Sprintf("%s - %s\n", op.Name, op.Doc))
		if len(op.DocExtra) > 0 {
			bw.WriteString("\n")
			bw.WriteString(op.DocExtra)
			bw.WriteString("\n")
		}
		bw.WriteString("*/\n")

		switch len(op.Imms) {
		case 0:
			fmt.Fprintf(bw, "%s avm_%s(%s);\n", rt, name, formatParams(op, 0))
		default:
			t := fmt.Sprintf("avm_%s_t", name)

			fmt.Fprintf(bw, "struct %s\n", t)
			bw.WriteString("{\n")
			for _, e := range op.Enum {
				if rt == "void" {
					bw.WriteString(fmt.Sprintf("\t%s %s(%s);\n", rt, e.Name, formatParams(op, 1)))
				} else {
					if len(op.Imms)+len(op.Stacks) > 1 {
						bw.WriteString(fmt.Sprintf("\t%s %s(%s);\n", e.Type, e.Name, formatParams(op, 1)))
					} else {
						bw.WriteString(fmt.Sprintf("\tconst %s %s;\n", e.Type, e.Name))
					}
				}
			}
			bw.WriteString("};\n")
			bw.WriteString(fmt.Sprintf("extern %s avm_%s;\n", t, name))
		}
	}

	return nil
}

func main() {
	var a args
	flag.StringVar(&a.Spec, "spec", "", "ceal spec file path")
	flag.StringVar(&a.Out, "out", "", "output file path")
	flag.Parse()

	err := run(a)
	if err != nil {
		panic(err)
	}

}
