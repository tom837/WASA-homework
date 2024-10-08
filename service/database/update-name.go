package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) UpdateName(name string, id string) error {
	query := "SELECT 1 FROM users WHERE username=? AND id!=?"
	err := db.c.QueryRow(query, name, id).Scan(new(int))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query = "UPDATE users SET username=? WHERE id=?"
			_, err := db.c.Exec(query, name, id)
			return err
		} else {
			return err
		}
	} else {
		return fmt.Errorf("Username already being used by a user")
	}

}
