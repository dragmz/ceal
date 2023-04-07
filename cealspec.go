package ceal

import (
	"fmt"
	"strings"
)

type CealArg struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type CealReturn struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type CealEnum struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type CealSpecOp struct {
	Name    string       `json:"name"`
	Imms    []CealArg    `json:"imms"`
	Stacks  []CealArg    `json:"stacks"`
	Returns []CealReturn `json:"returns"`
	Doc     string       `json:"doc"`
	Enum    []CealEnum   `json:"enum"`
}

type CealSpec struct {
	Version int          `json:"version"`
	Ops     []CealSpecOp `json:"ops"`
}

var replaceMap = map[string]string{
	"+":      "plus",
	"-":      "minus",
	"/":      "div",
	"*":      "mul",
	"<":      "lt",
	">":      "gt",
	"=":      "eq",
	"&":      "and",
	"|":      "or",
	"!":      "not",
	"%":      "mod",
	"^":      "xor",
	"~":      "inv",
	"switch": "switch_",
	"return": "return_",
}

func FormatOpName(name string) string {
	for from, to := range replaceMap {
		name = strings.ReplaceAll(name, from, to)
	}
	return name
}

func ReadReturnTypeName(op CealSpecOp) string {
	var rt string

	switch len(op.Returns) {
	case 0:
		rt = "void"
	case 1:
		rt = op.Returns[0].Type
	default:
		rt = fmt.Sprintf("avm_%s_result_t", op.Name)
	}

	return rt
}
