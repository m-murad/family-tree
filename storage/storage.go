package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type dataStore struct{ db *sql.DB }

type config struct {
	Url      string
	Port     string
	UserName string
	Password string
	Database string
}

func Connect() (Service, error) {
	config := config{
		Url:      "john.db.elephantsql.com",
		Port:     "5432",
		UserName: "gubgjfrw",
		Password: "RndaY1Y8STM5V7iQFpH7vfxjIHmLQ2v4",
		Database: "gubgjfrw",
	}

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		config.Url, config.Port, config.Database, config.UserName, config.Password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if err = createTables(db); err != nil {
		return nil, err
	}

	return dataStore{db: db}, nil
}

func createTables(db *sql.DB) error {

	createTables := `
		CREATE TABLE IF NOT EXISTS member (
			name VARCHAR(20) UNIQUE NOT NULL PRIMARY KEY
		);
		CREATE INDEX IF NOT EXISTS idx_member_name on member(name);

		CREATE TABLE IF NOT EXISTS relation (
			name VARCHAR(10) UNIQUE NOT NULL PRIMARY KEY
		);
		CREATE INDEX IF NOT EXISTS idx_relation_name on member(name);

		CREATE TABLE IF NOT EXISTS family_relation (
			first_member VARCHAR REFERENCES member(name),
			relation VARCHAR REFERENCES relation(name),
			second_member VARCHAR REFERENCES member(name),
			UNIQUE (first_member, relation, second_member),
			UNIQUE (first_member, second_member)
		);
		CREATE INDEX IF NOT EXISTS idx_relationship_first_name on family_relation(first_member);
		CREATE INDEX IF NOT EXISTS idx_relationship_second_name on family_relation(second_member);
`
	_, err := db.Exec(createTables)
	return err
}

func (d dataStore)Clear() error {
	statement := `TRUNCATE member, relation, relationship`
	_, err := d.db.Exec(statement)
	if err != nil {
		return errors.Errorf("Failed to delete data: %v", err)
	}
	return nil
}

type Service interface {
	Clear() error
	AddMember(member string) error
	AddRelation(relation string) error
	DefineRelation(firstMember string, relation string, secondMember string) error
	CountRelations(name string, relation string) (int, error)
	ListRelations(name string, relation string) (int, error)
}
