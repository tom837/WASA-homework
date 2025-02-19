package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) Comment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := rt.UserHandler(w, r, ps)
	photo := ps.ByName("id")
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	id, err := rt.db.AddComment(user, photo, comment.Comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := Comment{
		Id:      id,
		Userid:  user,
		Comment: comment.Comment}
	w.Header().Set("Content-Type", "encoding/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rt *_router) Comment_lst(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	photoid := ps.ByName("id")
	rows, err := rt.db.GetComments(photoid) // sql rows with all the users registered
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
	var user string
	var comment string
	var id string
	var data Comment
	lst := []Comment{} // list we will return with all the users
	defer rows.Close()
	for rows.Next() { // loop through all the users
		err = rows.Scan(&id, &user, &comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		username, err := rt.db.GetName(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data = Comment{ // create a json for the user
			Id:      id,
			Userid:  user,
			User:    username,
			Comment: comment,
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

func (rt *_router) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	user_id := rt.UserHandler(w, r, ps)
	err := rt.db.Remove_comment(id, user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Comment deleted successfully!")
}
