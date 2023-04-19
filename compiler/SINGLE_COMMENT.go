package compiler

import "strings"

func ceal_TrimSINGLE_COMMENTQuotes(v string) string {
	return strings.TrimPrefix(
		strings.Trim(v, "\r\n"),
		"//",
	)
}
