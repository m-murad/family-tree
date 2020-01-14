package storage

import (
	"github.com/pkg/errors"
)

func (d dataStore)AddRelation(relation string) error {
	statement := `INSERT INTO relation (name) VALUES ($1)`
	if _, err := d.db.Exec(statement, relation); err != nil {
		return errors.Errorf("Failed to add relation: %v", err)
	}
	return nil
}
