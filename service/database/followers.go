package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) Follow(user string, foll string, table string) error {
	if user == foll {
		return fmt.Errorf("cannot follow yourself")
	}
	query := "SELECT 1 FROM users WHERE id = ?;"
	err := db.c.QueryRow(query, foll).Scan(new(int)) // makes sure the id provided exists
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user not found")
		} else {
			return err
		}

	}
	query2 := fmt.Sprintf("SELECT follower_id FROM %s WHERE follower_id=? AND followed_id=?", table)
	var Id string
	err = db.c.QueryRow(query2, user, foll).Scan(&Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query3 := fmt.Sprintf("INSERT INTO %s (follower_id,followed_id) VALUES (?,?)", table)
			_, err1 := db.c.Exec(query3, user, foll)
			return err1
		} else {
			return err
		}
	} else if Id != "" {
		return fmt.Errorf("already following")

	} else {
		return err
	}

}

func (db *appdbimpl) GetFollowers(table string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s;", table)
	row, err := db.c.Query(query) // get all users in the database
	return row, err

}

func (db *appdbimpl) Following(id string) (*sql.Rows, error) {
	query := "SELECT followed_id FROM followers WHERE follower_id =? "
	row, err := db.c.Query(query, id)
	return row, err
}

func (db *appdbimpl) Unfollow(user string, fol string, table string) error {
	if user == fol {
		return fmt.Errorf("cannot unfollow yourself")
	}
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE follower_id=? AND followed_id=?", table)
	err := db.c.QueryRow(query, user, fol).Scan(new(int))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("you cannot unfollow someone you are not following")
		} else {
			return err
		}
	}
	query2 := fmt.Sprintf("DELETE FROM %s WHERE follower_id=? AND followed_id=?", table)
	_, err = db.c.Exec(query2, user, fol)
	return err
}
