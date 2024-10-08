package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) FollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	followed_id := ps.ByName("id")
	follower_id := rt.UserHandler(w, r, ps)
	err := rt.db.Follow(follower_id, followed_id, "followers")
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Followed successfuly!")
}

func (rt *_router) Follower_lst(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := rt.UserHandler(w, r, ps)
	rows, err := rt.db.Following(id)
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
	var following string
	lst := []string{} // list we will return with all the users
	defer rows.Close()
	for rows.Next() { // loop through all the users
		err = rows.Scan(&following)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		lst = append(lst, following) // add the json to the final list
	}
	w.Header().Set("Content-Type", "text/plain")
	err = json.NewEncoder(w).Encode(lst) // send the list of users
	if err != nil {
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rt *_router) Remove_follower(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	followed_id := ps.ByName("id")
	follower_id := rt.UserHandler(w, r, ps)
	err := rt.db.Unfollow(follower_id, followed_id, "followers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Unfollowed successfuly!")
}
