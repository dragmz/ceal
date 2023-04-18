package teal

import "github.com/dragmz/tealex"

type arguments struct {
	ts []tealex.Token
	i  int
}

func (a *arguments) Scan() bool {
	if a.i <= len(a.ts) {
		a.i++
		return a.i <= len(a.ts)
	}

	return false
}

func (a *arguments) Prev() tealex.Token {
	if a.i > 1 {
		return a.ts[a.i-2]
	}

	return tealex.Token{}
}

func (a *arguments) Curr() tealex.Token {
	if a.i > 0 && a.i <= len(a.ts) {
		return a.ts[a.i-1]
	}

	return tealex.Token{}
}

func (a *arguments) Text() string {
	return a.Curr().String()
}
