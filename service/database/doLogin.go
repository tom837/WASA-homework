package database

//doLogin as defined in api
import "database/sql"

func (db *appdbimpl) doLogin(name string) (string, error) {
	var id string
	err := db.c.QueryRow("SELECT id FROM users WHERE name=?;", name).Scan(&id) //get id from user with provided name
	if err == sql.ErrNoRows {                                                  //if no user hase that name create it
		err = db.SetName(name)
		if err != nil { //if there is an error return it
			return "", err
		}
		err = db.c.QueryRow("SELECT id FROM users WHERE name=?;", name).Scan(&id)
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
