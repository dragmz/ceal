package compiler

import "strings"

func ceal_TrimSTRINGQuotes(v string) string {
	return strings.TrimSuffix(strings.TrimPrefix(v, "\""), "\"")
}
