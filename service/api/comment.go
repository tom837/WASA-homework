package api

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)



func (rt *_router) Comment(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	user:= rt.UserHandler(w,r,ps)
	photo := ps.ByName("id")
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err = rt.db.AddComment(user, photo, comment.Comment)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w,"Comment added successfully!")

}

func (rt *_router) Comment_lst(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	photoid := ps.ByName("id")
	fmt.Println(photoid)
	rows,err := rt.db.GetComments(photoid) //sql rows with all the users registered
	var user string
	var comment string
	var id string
	var data Comment
	lst := []Comment{}  //list we will return with all the users
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	for rows.Next() { //loop through all the users
		err = rows.Scan(&id,&user,&comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data = Comment{  //create a json for the user
			Id : id,
			User: user,
			Comment: comment,
		}
		lst = append(lst, data) //add the json to the final list
	}
	w.Header().Set("Content-Type", "text/plain")
	err = json.NewEncoder(w).Encode(lst) //send the list of users
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



func (rt *_router) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	id := ps.ByName("id")
	user_id:=rt.UserHandler(w,r,ps)
	err := rt.db.Remove_comment(id, user_id)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w,"Comment deleted successfully!")
}