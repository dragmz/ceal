package teal

import "strings"

type Source struct {
	lines []string
}

func (l *Source) WriteLine(s string) {
	if s == "" {
		return
	}

	l.lines = append(l.lines, s)
}

func (l *Source) String() string {
	return strings.Join(l.lines, "\n")
}
