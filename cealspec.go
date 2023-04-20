package ceal

import (
	"fmt"
	"regexp"
	"strings"
)

type CealArg struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Array bool   `json:"array"`
	Field bool   `json:"field"`
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
	Name     string       `json:"name"`
	Imms     []CealArg    `json:"imms"`
	Stacks   []CealArg    `json:"stacks"`
	Returns  []CealReturn `json:"returns"`
	Doc      string       `json:"doc"`
	DocExtra string       `json:"docExtra"`
	Enum     []CealEnum   `json:"enum"`
}

type CealSpecType struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
}

type CealSpec struct {
	Version int            `json:"version"`
	Types   []CealSpecType `json:"types"`
	Ops     []CealSpecOp   `json:"ops"`
}

var replaceMap = map[string]string{
	"+": "plus",
	"-": "minus",
	"/": "div",
	"*": "mul",
	"<": "lt",
	">": "gt",
	"=": "eq",
	"&": "and",
	"|": "or",
	"!": "not",
	"%": "mod",
	"^": "xor",
	"~": "inv",
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

var ArrayRegex = regexp.MustCompile(`^\[(\d+?)\]`)

func ReadStackType(name string) string {
	arr := ArrayRegex.FindStringSubmatch(name)
	if arr != nil {
		name = strings.ReplaceAll(name, arr[0], fmt.Sprintf("r%s_", arr[1]))
	}

	name = strings.ReplaceAll(name, "[]", "r_")
	name = fmt.Sprintf("%s_t", name)
	return name
}
