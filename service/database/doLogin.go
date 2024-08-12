package database

//doLogin as defined in api
import "database/sql"

func (db *appdbimpl) DoLogin(name string) (string, error) {
	query1 := "SELECT id FROM users WHERE username = ?"
	var id string
	err := db.c.QueryRow(query1, name).Scan(&id) //get id from user with provided name
	if err == sql.ErrNoRows {                    //if no user hase that name create it
		id, err = db.SetName(name)
		if err != nil { //if there is an error return it
			return "", err
		}
		return id, err
	} else if err != nil { //if there is an error return it
		return "", err
	} else {
		return id, err
	}

}
