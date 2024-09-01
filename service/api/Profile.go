package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"time"
	"database/sql"
	"fmt"
	"sort"
	"log"
)



func (rt *_router) UserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user string)([]byte){
	var name string
	rows,name,err := rt.db.GetProfile(user) //sql rows with all the users registered
	var comment sql.NullString
	var id string
	var photo []byte
	var like sql.NullString
	var time time.Time
	var data Post
	lst := []Post{}  //list we will return with all the users
	if err != nil {
		var nothing []byte
		// Handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nothing
	}
	defer rows.Close()
	for rows.Next() { //loop through all the users
		err = rows.Scan(&photo,&id,&like,&comment,&time)
		if err != nil {
			var nothing []byte
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nothing
		}
		data = Post{  //create a json for the user
			User: 		name,
			Id: 		id,
			Photo: 		photo,
			Likes:		like.String,
			Comments: 	comment.String,
			Time: 		time,
		}
		lst = append(lst, data) //add the json to the final list
	}
	// Map to hold aggregated data keyed by user and photo
	aggregated := make(map[string]*AggregatedData)

	for _, entry := range lst {
		key := entry.User + "-" + entry.Id
		if _, exists := aggregated[key]; !exists {
			aggregated[key] = &AggregatedData{
				User:     entry.User,
				Id:       entry.Id,
				Photo:    entry.Photo,
				Likes:    []string{},
				Comments: []string{},
				Time:     entry.Time,
			}
		}

		// Add unique likes
		if !contains(aggregated[key].Likes, entry.Likes) {
			aggregated[key].Likes = append(aggregated[key].Likes, entry.Likes)
		}

		// Add unique comments
		if !contains(aggregated[key].Comments, entry.Comments) {
			aggregated[key].Comments = append(aggregated[key].Comments, entry.Comments)
		}
	}

	// Convert the map to a slice
	var result []AggregatedData
	for _, data := range aggregated {
		result = append(result, *data)
	}

   	// Convert the result to JSON
   	jsonData, err := json.Marshal(result)
   	if err != nil {
		var nothing []byte
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nothing
   	}
	return jsonData
}

// Helper function to check if a slice contains a value
func contains(slice []string, value string) bool {
   for _, v := range slice {
	   if v == value {
		   return true
	   }
   }
   return false
}

func (rt *_router) Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params)(){
	user:=ps.ByName("id")
	profile := rt.UserProfile(w,r,ps,user)
	if string(profile) == "null"{
		fmt.Fprintf(w,"No posts yet")
	}else{
		fmt.Fprintf(w,string(profile))
	}
}




func (rt *_router) Stream(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	user:=rt.UserHandler(w,r,ps)
	rows,err := rt.db.GetStream(user)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	lst:=[]string{}
	var id string
	for rows.Next() { //loop through all the users
		err = rows.Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		lst = append(lst, id) //add the json to the final list
	}
	var other []byte
	for _, value := range lst {
		tmp:=rt.UserProfile(w,r,ps,value)
		if string(tmp)!="null"{
			var array1 []map[string]interface{}
			var array2 []map[string]interface{}

			// Unmarshal both JSON arrays into Go slices
			if len(other)>0 {
				if err := json.Unmarshal(other, &array1); err != nil {
					log.Fatal(err)
				}
			}
			if err := json.Unmarshal(tmp, &array2); err != nil {
				log.Fatal(err)
			}
		
			// Combine the slices
			combinedjsons := append(array1, array2...)
		
			// Marshal the combined slice back into JSON
			other, err = json.Marshal(combinedjsons)
			if err != nil {
				log.Fatal(err)
			}
		}
    }
	if other !=nil{
		var photos []AggregatedData
		err = json.Unmarshal(other, &photos)
		if err != nil {
			log.Fatal(err)
		}

		// Sort the photos slice by the Time field
		sort.Slice(photos, func(i, j int) bool {
			return photos[i].Time.After(photos[j].Time)
		})

		// Marshal the sorted slice back into JSON
		sortedJSON, err := json.MarshalIndent(photos, "", "    ")
		if err != nil {
			log.Fatal(err)
		}


		fmt.Fprintf(w,string(sortedJSON))
	}
	
	

}