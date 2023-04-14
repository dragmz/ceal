package compiler

import "strconv"

func itoa(v int) string {
	return strconv.Itoa(v)
}

func atoi(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}
