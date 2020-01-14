package storage

import (
	"github.com/pkg/errors"
)

func (d dataStore)AddMember(memberName string) error {
	statement := `INSERT INTO member (name) VALUES ($1)`
	if _, err := d.db.Exec(statement, memberName); err != nil {
		return errors.Errorf("Failed to add member: %v", err)
	}
	return nil
}
