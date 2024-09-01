package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)


func (rt *_router) UserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)(string) {
	// Extract user ID from the session
	userID:=r.Header.Get("Authorization")
	return userID
}


func (rt *_router) Whoami(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	userID:=rt.UserHandler(w,r,ps)
	fmt.Fprintf(w,userID)
}
