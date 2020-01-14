package storage

import (
	"github.com/pkg/errors"
	"strings"
)

func (d dataStore)DefineRelation(firstMember string, relation string, secondMember string) error {
	statement := `INSERT INTO family_relation (first_member, relation, second_member) VALUES ($1, $2, $3)`
	if _, err := d.db.Exec(statement, firstMember, relation, secondMember); err != nil {
		if strings.Contains(err.Error(), `violates foreign key constraint "family_relation_relation_fkey"`) {
			return errors.Errorf("%s is not a relation", relation)
		}
		if strings.Contains(err.Error(), `violates foreign key constraint "family_relation_first_member_fkey"`) {
			return errors.Errorf("%s is not a member", firstMember)
		}
		if strings.Contains(err.Error(), `violates foreign key constraint "family_relation_second_member_fkey"`) {
			return errors.Errorf("%s is not a member", secondMember)
		}
		if strings.Contains(err.Error(), `violates unique constraint "family_relation_first_member_relation_second_member_key"`) {
			return errors.New("This relationship already exists")
		}
		if strings.Contains(err.Error(), `violates unique constraint "family_relation_first_member_second_member_key"`) {
			return errors.Errorf("A different relationship already exists between %s and %s", firstMember, secondMember)
		}

		return errors.Errorf("Failed to add member: %v", err)
	}
	return nil
}

func (d dataStore)CountRelations(name string, relation string) (int, error) {
	var count int

	qry := `SELECT COUNT(*) FROM family_relation WHERE first_member = $1 AND relation = $2`
	if err := d.db.QueryRow(qry, name, relation).Scan(&count); err != nil {
		return 0, errors.Errorf("Failed to count relationships: %v", err)
	}

	return count, nil
}

func (d dataStore)ListRelations(name string, relation string) ([]string, error) {
	var relations []string

	qry := `SELECT second_member FROM family_relation WHERE first_member = $1 AND relation = $2`
	rows, err := d.db.Query(qry, name, relation)
	if err != nil {
		return nil, errors.Errorf("Failed to list relationships: %v", err)
	}

	for rows.Next() {
		var relation string

		if err = rows.Scan(&relation); err != nil {
			return nil, errors.Errorf("Failed to list relationships: %v", err)
		}
		relations = append(relations, relation)
	}

	return relations, nil
}
