package database

import (
	"database/sql"
	"fmt"
)


func (db *appdbimpl) GetProfile(user string)(*sql.Rows, string, error){
	query:= "SELECT username FROM users WHERE id=?"
	row, err:= db.c.Query(query, user)
	if err!=nil{
		var nothing string
		return row,nothing,err
	}
	var name string
	for row.Next() { //loop through all the users
		err = row.Scan(&name)
		if err != nil {
			return row,name, err
		}
	}
	if name == ""{
		return row,name, fmt.Errorf("User not found")
	}
	query = "SELECT photos.photo, photos.id, likes.user_id, comments.comment, uploaded_at  FROM photos LEFT JOIN likes ON photos.id=likes.photo_id LEFT JOIN comments ON photos.id=comments.photo_id WHERE photos.user_id=?;"
	//query := "SELECT * FROM photos JOIN likes ON photos.id=likes.photo_id WHERE photos.user_id=?;"
	row, err= db.c.Query(query, user) // get all photos in the database
	return row, name,err
}

func (db *appdbimpl) GetStream(user string)(*sql.Rows, error){
	query:= "SELECT id FROM users WHERE id=?"
	row, err:= db.c.Query(query, user)
	if err!=nil{
		return row,err
	}
	var id string
	for row.Next() { //loop through all the users
		err = row.Scan(&id)
		if err != nil {
			return row, err
		}
	}
	if id == ""{
		return row, fmt.Errorf("User not found")
	}
	query = "SELECT followed_id FROM followers WHERE follower_id=?"
	row, err= db.c.Query(query, user) // get all photos in the database
	return row,err
	
}