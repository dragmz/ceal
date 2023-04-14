package compiler

import "strings"

type Lines struct {
	lines []string
}

func (l *Lines) WriteLine(s string) {
	if s == "" {
		return
	}

	l.lines = append(l.lines, s)
}

func (l *Lines) String() string {
	return strings.Join(l.lines, "\n")
}
