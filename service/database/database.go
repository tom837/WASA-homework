package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName(id string) (string, error)
	SetName(name string) (string, error)
	DoLogin(name string) (string, error)
	Getusers()(*sql.Rows, error)
	UpdateName(name string, id string) (error)
	Ping() error
	Follow(user string, foll string, table string)(error)
	GetFollowers(table string)(*sql.Rows, error)
	Unfollow(user string, fol string, table string)(error)
	Upload(user string, photo []byte) (error)
	GetPhotos()(*sql.Rows, error)
	Remove_photo(id string, user_id string)(error)
	LikePhoto(user string, photo string) (error)
	GetLikes()(*sql.Rows, error)
	Unlike(user string, photo string)(error)
	AddComment(user string, photo string, comment string) (string, error)
	GetComments(photoid string)(*sql.Rows, error)
	Remove_comment(id string, user_id string)(error)
	GetProfile(user string)(*sql.Rows, string, error)
	GetStream(user string)(*sql.Rows, error)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	tabel_users := "CREATE TABLE IF NOT EXISTS users(id INT NOT NULL PRIMARY KEY,username TEXT NOT NULL);"
	Follower_table := "CREATE TABLE IF NOT EXISTS followers(follower_id INT NOT NULL, followed_id INT NOT NULL, PRIMARY KEY (follower_id, followed_id), FOREIGN KEY (follower_id) REFERENCES users(id), FOREIGN KEY (followed_id) REFERENCES users(id));"
	Banned_table :="CREATE TABLE IF NOT EXISTS baned(follower_id INT NOT NULL, followed_id INT NOT NULL, PRIMARY KEY (follower_id, followed_id), FOREIGN KEY (follower_id) REFERENCES users(id), FOREIGN KEY (followed_id) REFERENCES users(id));"
	Photo_table := "CREATE tABLE IF NOT EXISTS photos(id INT NOT NULL PRIMARY KEY,user_id INT NOT NULL, photo BYTEA NOT NULL, uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (user_id) REFERENCES users(id));"
	Like_table :="CREATE TABLE IF NOT EXISTS likes(user_id INT NOT NULL, photo_id INT NOT NULL, PRIMARY KEY (user_id, photo_id), FOREIGN KEY (user_id) REFERENCES users(id), FOREIGN KEY (photo_id) REFERENCES photos(id));"
	Comment_table := "CREATE tABLE IF NOT EXISTS comments(id INT NOT NULL PRIMARY KEY,user_id INT NOT NULL,photo_id INT NOT NULL, comment TEXT NOT NULL , FOREIGN KEY (user_id) REFERENCES users(id),FOREIGN KEY (photo_id) REFERENCES photos(id));"
	_, err := db.Exec(tabel_users)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(Follower_table)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(Banned_table)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(Photo_table)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(Like_table)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(Comment_table)
	if err != nil {
		return nil, err
	}
	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{c: db}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
