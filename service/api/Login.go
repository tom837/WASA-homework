package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/sessions"
	"fmt"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func (rt *_router) AssertNameCorrect(name string) (bool) {
	if len(name)<3 || len(name)>16{
		return true
	}
	return false

}

func (rt *_router) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	id, err := rt.db.DoLogin(user.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	//Store id into session
	session, _ := store.Get(r, "session-name")
	session.Values["userID"] = id
	session.Save(r, w)
	fmt.Fprintf(w,"You logend in successfully!")
	
	//Make usename and id into User structure
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
