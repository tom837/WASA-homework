package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"time"
	"io"

)

func (rt *_router) UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	fmt.Println("loading")
	user_id :=rt.UserHandler(w,r,ps)
	photo, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err = rt.db.Upload(user_id, photo)
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Photo uploaded successfuly!")

}

func (rt *_router) Photo_lst(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	rows,err := rt.db.GetPhotos() //sql rows with all the users registered
	var user string
	var id string
	var time time.Time
	var data Photo
	lst := []Photo{}  //list we will return with all the users
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	for rows.Next() { //loop through all the users
		err = rows.Scan(&id,&user,&time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data = Photo{  //create a json for the user
			Id: id,
			User: user,
			Time: time,
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


func (rt *_router) DeletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	id := ps.ByName("id")
	user_id:=rt.UserHandler(w,r,ps)
	err := rt.db.Remove_photo(id, user_id)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w,"Photo deleted successfully!")
}