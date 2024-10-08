package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) Ban(user string, foll string) error {
	if user == foll {
		return fmt.Errorf("Cannot ban yourself!")
	}
	query := "SELECT 1 FROM users WHERE id = ?;"
	err := db.c.QueryRow(query, foll).Scan(new(int)) // makes sure the id provided exists
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("User not found!")
		} else {
			return err
		}

	}
	query2 := "SELECT follower_id FROM followers WHERE follower_id=? AND followed_id=?"
	var Id string
	err = db.c.QueryRow(query2, user, foll).Scan(&Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query3 := "INSERT INTO followers (follower_id,followed_id) VALUES (?,?)"
			_, err1 := db.c.Exec(query3, user, foll)
			return err1
		} else {
			return err
		}
	} else if Id != "" {
		return fmt.Errorf("Already banned!")

	} else {
		return err
	}

}

func (db *appdbimpl) GetBanned(id string) (*sql.Rows, error) {
	query := "SELECT followed_id FROM baned WHERE follower_id=?;"
	row, err := db.c.Query(query, id)
	return row, err
}

func (db *appdbimpl) Unban(user string, fol string) error {
	if user == fol {
		return fmt.Errorf("Cannot unban yourself!")
	}
	query := "SELECT 1 FROM followers WHERE follower_id=? AND followed_id=?"
	err := db.c.QueryRow(query, user, fol).Scan(new(int))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("You cannot unban someone you did not ban!")
		} else {
			return err
		}
	}
	query2 := "DELETE FROM followers WHERE follower_id=? AND followed_id=?"
	_, err = db.c.Exec(query2, user, fol)
	return err

}
