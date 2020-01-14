package cmd

import (
	"github.com/pkg/errors"
	"strings"
)

func validName(name string) (string, error) {
	name = strings.TrimSpace(strings.ToUpper(name))
	if len(name) < 3 || len(name) > 20 {
		return "", errors.New("Member name should be 3 - 20 character long.")
	}
	return name, nil
}

