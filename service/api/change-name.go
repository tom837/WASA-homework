package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) SetMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if rt.AssertNameCorrect(user.UserName) {
		err = fmt.Errorf("Username has to be between 3 and 16 characters long")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := rt.UserHandler(w, r, ps)
	err = rt.db.UpdateName(user.UserName, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := rt.db.DoLogin(user.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if id != userID {
		err = fmt.Errorf("id missmatch")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		return
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
		return
	}
}
