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

func validRelation(relation string) (string, error) {
	relation = strings.TrimSpace(strings.ToUpper(relation))
	if len(relation) < 3 || len(relation) > 10 {
		return "", errors.New("Relation name should be 3 - 10 character long.")
	}
	return relation, nil
}
