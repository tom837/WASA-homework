package api

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func (rt *_router) Like(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	user:= rt.UserHandler(w,r,ps)
	photo:= ps.ByName("id")
	err := rt.db.LikePhoto(user, photo)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w,"Photo liked successfully!")
}


func (rt *_router) Like_lst(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	rows,err := rt.db.GetLikes() //sql rows with all the users registered
	var user string
	var photo string
	var data Like
	lst := []Like{}  //list we will return with all the users
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	for rows.Next() { //loop through all the users
		err = rows.Scan(&user,&photo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data = Like{  //create a json for the user
			User: user,
			Photo: photo,
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


func (rt *_router) DeleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	user := rt.UserHandler(w,r,ps)
	photo:= ps.ByName("id")
	err := rt.db.Unlike(user,photo)
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	fmt.Fprintf(w,"unliked successfully!")
}