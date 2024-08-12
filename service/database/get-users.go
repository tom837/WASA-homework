package database

//doLogin as defined in api
import "database/sql"

func (db *appdbimpl) Getusers() (*sql.Rows, error) {
	query := "SELECT * FROM users;"
	//var id string
	//res,err := db.c.Exec(query)
	row, err:= db.c.Query(query) // get all users in the database
	return row,err
	
}


