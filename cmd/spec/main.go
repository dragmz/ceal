package main

import (
	"ceal"
	"encoding/json"
	"flag"
	"fmt"
	"os"

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
	case ".":
		return "any"
	default:
		return t
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
	imms := op.Size - 1

	ps := []ceal.CealArg{}

	for i := 0; i < imms; i++ {
		t := "uint64"

		p := ceal.CealArg{
			Type: t,
			Name: fmt.Sprintf("i%d", i+1),
		}

		ps = append(ps, p)
	}

	return ps
}

func readStacks(op LangSpecOp) []ceal.CealArg {
	stacks := len(op.Args)

	ps := []ceal.CealArg{}

	for i := 0; i < stacks; i++ {
		p := ceal.CealArg{
			Type: readType(op.Args[i]),
			Name: fmt.Sprintf("s%d", i+1),
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
