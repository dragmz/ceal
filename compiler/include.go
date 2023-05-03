package compiler

import (
	"fmt"
	"os"
	"path"

	"github.com/pkg/errors"
)

type cealIncludes struct {
	paths []string
}

func (i *cealIncludes) resolve(quoted string) (string, string, error) {
	begin := string(quoted[0])
	var end string

	switch begin {
	case "\"":
		end = begin
	case "<":
		end = ">"
	default:
		return "", "", fmt.Errorf("unexpected #include quote in: '%s'", quoted)
	}

	if end != string(quoted[len(quoted)-1]) {
		return "", "", fmt.Errorf("#include quote mismatch: '%s'", quoted)
	}

	name := quoted[1 : len(quoted)-1]

	for _, dir := range i.paths {
		ip := path.Join(dir, name)

		bs, err := os.ReadFile(ip)
		if os.IsNotExist(err) {
			continue
		}

		if err != nil {
			return "", "", errors.Wrapf(err, "failed to read #include file: '%s'", name)
		}

		inc := string(bs)
		return inc, ip, nil
	}

	return "", "", errors.Errorf("failed to find #include file: '%s'", name)
}
