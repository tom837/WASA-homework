package database


import (
	"fmt"
	"database/sql"
)


func (db *appdbimpl) Upload(user string, photo []byte) (error){
	newKey, err := generateNewKey(db, "p", "photos")
	query := `INSERT INTO photos (id,user_id, photo) VALUES (?, ?, ?);`
	_, err = db.c.Exec(query,newKey, user, photo)
	return err
}


func (db *appdbimpl) GetPhotos()(*sql.Rows, error){
	query := "SELECT id,user_id,uploaded_at  FROM photos;"
	row, err:= db.c.Query(query) // get all photos in the database
	return row,err
	
}


func (db *appdbimpl) Remove_photo(id string, user_id string)(error){
	var user string
	query := "SELECT user_id FROM photos WHERE id=?"
	err := db.c.QueryRow(query, id).Scan(&user)
	if err!= nil{
		if err == sql.ErrNoRows{
			return fmt.Errorf("Photo not found")
		}else{
			return err
		}
	}
	if user != user_id{
		return fmt.Errorf("You cannot delete another users photo!")
	}
	query = "DELETE FROM photos WHERE id=?"
	_, err= db.c.Exec(query, id)
	if err!=nil{
		return err
	}
	query = "DELETE FROM likes WHERE photo_id=?"
	_, err= db.c.Exec(query, id)
	if err!=nil{
		return err
	}
	query = "DELETE FROM comments WHERE photo_id=?"
	_, err= db.c.Exec(query, id)
	return err
}