package compiler

import "strings"

func ceal_TrimMULTILINE_COMMENTQuotes(v string) string {
	return strings.TrimSuffix(strings.TrimPrefix(v, "/*"), "*/")
}
