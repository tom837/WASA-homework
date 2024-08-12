package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)


func (rt *_router) SetMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if rt.AssertNameCorrect(user.UserName){
		fmt.Fprintf(w,"Username hase to be between 3 and 16 characters long")
		return
	}
	userID :=rt.UserHandler(w,r,ps)
	err = rt.db.UpdateName(user.UserName, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id, err := rt.db.DoLogin(user.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if id!= userID{
		w.Header().Set("content-type", "text/plain")
		_, _ = w.Write([]byte("id missmatch"))
	}
	data := User{
		UserName: user.UserName,
		ID:       id,
	}
	w.Header().Set("Content-Type", "encoding/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	




}