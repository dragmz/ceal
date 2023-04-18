package main

import (
	"ceal"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type LangSpec struct {
	EvalMaxVersion int          `json:"EvalMaxVersion"`
	Ops            []LangSpecOp `json:"ops"`
}

type OpImmediateDetails struct {
	Comment   string `json:",omitempty"`
	Encoding  string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Reference string `json:",omitempty"`
}

type LangSpecOp struct {
	Opcode  byte
	Name    string
	Args    []string `json:",omitempty"`
	Returns []string `json:",omitempty"`
	Size    int

	ArgEnum      []string `json:",omitempty"`
	ArgEnumTypes []string `json:",omitempty"`

	Doc               string
	DocExtra          string               `json:",omitempty"`
	ImmediateNote     []OpImmediateDetails `json:",omitempty"`
	IntroducedVersion uint64
	Groups            []string
}

type args struct {
	Spec string
	Out  string
}

func readType(t string) string {
	switch t {
	case "[]byte", "addr", "[32]byte", "key", "bigint":
		return "bytes"
	case "uint64", "bool":
		return "uint64"
	case "any":
		return "any"
	default:
		panic(fmt.Sprintf("unsupported type: '%s'", t))
	}
}

func readEnum(op LangSpecOp) []ceal.CealEnum {
	es := []ceal.CealEnum{}

	for i, name := range op.ArgEnum {
		var t string

		switch len(op.ArgEnumTypes) {
		case 0:
			t = "void"
		default:
			t = readType(op.ArgEnumTypes[i])
		}

		e := ceal.CealEnum{
			Type: t,
			Name: name,
		}

		es = append(es, e)
	}

	return es
}

func readReturns(op LangSpecOp) []ceal.CealReturn {
	rs := []ceal.CealReturn{}

	for i := range op.Returns {
		rt := readType(op.Returns[i])
		rs = append(rs, ceal.CealReturn{
			Type: rt,
			Name: fmt.Sprintf("r%d", i+1),
		})
	}

	return rs
}

func readImms(op LangSpecOp) []ceal.CealArg {
	imms := len(op.ImmediateNote)

	ps := make([]ceal.CealArg, imms)
	if imms == 0 {
		return ps
	}

	for i := 0; i < imms; i++ {
		imm := op.ImmediateNote[i]

		t := "uint64"
		array := false
		switch imm.Encoding {
		case "uint8":
			t = "uint8"
		case "int8":
			t = "int8"
		case "int16 (big-endian)":
			t = "int16"
		case "varuint":
			t = "bytes"
		case "varuint count, [int16 (big-endian) ...]":
			t = "bytes"
			array = true
		case "varuint count, [varuint ...]":
			t = "bytes"
			array = true
		case "varuint count, [varuint length, bytes ...]":
			t = "bytes"
			array = true
		case "varuint length, bytes":
			t = "bytes"
		default:
			panic(fmt.Sprintf("Unsupported immediate type: %s", imm.Encoding))
		}

		name := fmt.Sprintf("%s%d", strings.Trim(imm.Name, " ."), i+1)
		ps[i] = ceal.CealArg{
			Type:  t,
			Name:  name,
			Array: array,
		}
	}
	return ps
}

func readStacks(op LangSpecOp) []ceal.CealArg {
	stacks := len(op.Args)

	ps := []ceal.CealArg{}

	for i := 0; i < stacks; i++ {
		p := ceal.CealArg{
			Type: readType(op.Args[i]),
			Name: fmt.Sprintf("STACK_%d", i+1),
		}

		ps = append(ps, p)
	}

	return ps
}

func run(a args) error {
	r, err := os.Open(a.Spec)
	if err != nil {
		return errors.Wrap(err, "failed to download langspec")
	}

	defer r.Close()

	var ls LangSpec

	d := json.NewDecoder(r)
	err = d.Decode(&ls)
	if err != nil {
		return errors.Wrap(err, "failed to decode langspec")
	}

	cs := ceal.CealSpec{
		Version: ls.EvalMaxVersion,
	}

	for _, op := range ls.Ops {
		cop := ceal.CealSpecOp{
			Name:     op.Name,
			Doc:      op.Doc,
			DocExtra: op.DocExtra,
			Imms:     readImms(op),
			Stacks:   readStacks(op),
			Returns:  readReturns(op),
			Enum:     readEnum(op),
		}

		cs.Ops = append(cs.Ops, cop)
	}

	f, err := os.Create(a.Out)
	if err != nil {
		return errors.Wrap(err, "failed to create lanspec output file")
	}

	defer f.Close()

	e := json.NewEncoder(f)
	e.SetIndent("", "\t")

	e.Encode(cs)

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
