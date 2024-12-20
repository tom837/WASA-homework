package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) Users_lst(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.Getusers() // sql rows with all the users registered
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rows.Err() != nil {
		rows.Close()
		http.Error(w, rows.Err().Error(), http.StatusInternalServerError)
		return
	}
	var id string
	var name string
	var data User
	lst := []User{} // list we will return with all the users
	defer rows.Close()
	for rows.Next() { // loop through all the users
		err := rows.Scan(&name, &id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data = User{ // create a json for the user
			UserName: id,
			ID:       name,
		}
		lst = append(lst, data) // add the json to the final list
	}
	w.Header().Set("Content-Type", "text/plain")
	err = json.NewEncoder(w).Encode(lst) // send the list of users
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
