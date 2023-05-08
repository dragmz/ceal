package compiler

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type cealIncludes struct {
	paths []string
}

func (i *cealIncludes) resolve(name string, global bool) (string, []byte, error) {
	for _, dir := range i.paths {
		ip := filepath.Join(dir, name)

		id, err := filepath.Abs(ip)
		if err != nil {
			return "", nil, errors.Wrapf(err, "failed to resolve #include file: '%s'", name)
		}

		bs, err := os.ReadFile(ip)
		if os.IsNotExist(err) {
			continue
		}

		if err != nil {
			return "", nil, errors.Wrapf(err, "failed to read #include file: '%s'", name)
		}

		return id, bs, nil
	}

	return "", nil, errors.Errorf("failed to find #include file: '%s'", name)
}
