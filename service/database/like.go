package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) LikePhoto(user string, photo string) error {
	query := "SELECT 1 FROM photos WHERE id=?"
	err := db.c.QueryRow(query, photo).Scan(new(int))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("Photo not found")
		} else {
			return err
		}
	}
	query = "SELECT 1 FROM likes WHERE user_id=? AND photo_id=?"
	err = db.c.QueryRow(query, user, photo).Scan(new(int))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query = "INSERT INTO likes (user_id,photo_id) VALUES (?,?);"
			_, err = db.c.Exec(query, user, photo)
			return err
		} else {
			return err
		}
	} else {
		return fmt.Errorf("You cannot like a photo twice!")
	}

}

func (db *appdbimpl) GetLikes() (*sql.Rows, error) {
	query := "SELECT *  FROM likes;"
	row, err := db.c.Query(query) // get all likes in the database
	return row, err

}

func (db *appdbimpl) Unlike(user string, photo string) error {
	query := "SELECT user_id FROM photos WHERE id=?"
	err := db.c.QueryRow(query, photo).Scan(new(int))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("Photo not found!")
		} else {
			return err
		}
	}

	query = "SELECT 1 FROM likes WHERE photo_id=? AND user_id=?"
	err = db.c.QueryRow(query, photo, user).Scan(new(int))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("Cannot unlike a photo you haven't liked!")
		} else {
			return err
		}
	}
	query = "DELETE FROM likes WHERE user_id=? AND photo_id=?"
	_, err = db.c.Exec(query, user, photo)
	return err
}
