package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)


func (rt *_router) UserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)(string) {
	// Extract user ID from the session
	session, _ := store.Get(r, "session-name")
	ID, ok := session.Values["userID"]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return ""
	}
	userID, ok := ID.(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return ""
	}
	return userID
	
}

func (rt *_router) AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session, _ := store.Get(r, "session-name")
		if _, ok := session.Values["userID"]; !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r, ps)
	}
		
}


func (rt *_router) Whoami(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	userID:=rt.UserHandler(w,r,ps)
	fmt.Fprintf(w,userID)
}
