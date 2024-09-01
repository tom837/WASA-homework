package database


import (
	"fmt"
	"database/sql"
)


func (db *appdbimpl) AddComment(user string, photo string, comment string) (string, error){
	query:="SELECT id FROM photos WHERE id=?"
	err := db.c.QueryRow(query, photo).Scan(&photo)
	if err != nil{
		if err == sql.ErrNoRows{
			return "", fmt.Errorf("Photo not found")
		}else{
			return "", err
		}
	}
	newKey, err := generateNewKey(db, "c", "comments")
	query = `INSERT INTO comments (id,user_id, photo_id,comment) VALUES (?, ?, ?, ?);`
	_, err = db.c.Exec(query,newKey, user, photo, comment)
	return newKey, err
}



func (db *appdbimpl) GetComments(photoid string)(*sql.Rows, error){
	query := "SELECT id,user_id,comment  FROM comments WHERE photo_id=?;"
	row, err:= db.c.Query(query,photoid)
	return row,err
}




func (db *appdbimpl) Remove_comment(id string, user_id string)(error){
	var user string
	query := "SELECT user_id FROM comments WHERE id=?"
	err := db.c.QueryRow(query, id).Scan(&user)
	if err!= nil{
		if err == sql.ErrNoRows{
			return fmt.Errorf("Comment not found")
		}else{
			return err
		}
	}
	if user != user_id{
		return fmt.Errorf("You cannot delete another users comment!")
	}
	query = "DELETE FROM comments WHERE id=?"
	_, err= db.c.Exec(query, id)
	return err
}