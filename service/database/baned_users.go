package database


import (
	"fmt"
	"database/sql"
)

func (db *appdbimpl) Ban(user string, foll string) (error){
	if user == foll {
		return fmt.Errorf("Cannot follow yourself!")
	}
	query := "SELECT id FROM users WHERE id = ?;"
	var Identifier string
	err := db.c.QueryRow(query, foll).Scan(&Identifier) //makes sure the id provided exists
	if err != nil {
		eror := err.Error()
		if eror == "sql: no rows in result set"{
			return fmt.Errorf("User not found!")
		}else{
			return err
		}
		
	}
	query2:="SELECT follower_id FROM followers WHERE follower_id=? AND followed_id=?"
	var Id string
	err = db.c.QueryRow(query2, user, foll).Scan(&Id)
	
	if err != nil{
		eror:= err.Error()
		if eror== "sql: no rows in result set"{
			query3 := "INSERT INTO followers (follower_id,followed_id) VALUES (?,?)"
			_, err1 := db.c.Exec(query3, user, foll)
			return err1
		}else{
			return err
		}
	} else if Id!="" {
		return fmt.Errorf("Already following!")
	
	} else {
		return err
	}
	
}


func (db *appdbimpl) GetBanned(id string)(*sql.Rows, error){
	query := "SELECT followed_id FROM baned WHERE follower_id=?;"
	row, err:= db.c.Query(query,id)
	return row,err
}


func (db *appdbimpl) Unban(user string, fol string)(error){
	if user == fol {
		return fmt.Errorf("Cannot unfollow yourself!")
	}
	query:="SELECT follower_id FROM followers WHERE follower_id=? AND followed_id=?"
	var Id string
	err := db.c.QueryRow(query, user, fol).Scan(&Id)
	if err!=nil{
		eror:= err.Error()
		if eror== "sql: no rows in result set"{
			return fmt.Errorf("You cannot unfollow someone you are not following!")
		}else{
			return err
		}
		
	}

	query2 := "DELETE FROM followers WHERE follower_id=? AND followed_id=?"
	_, err = db.c.Exec(query2, user, fol)
	return err
	
}