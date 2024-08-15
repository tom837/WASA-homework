package api

import (
	"time"
)

type User struct {
	UserName string "json:\"name\""
	ID       string "json:\"ID\""
	//Email    string "json:\"email\""
}


type Follow struct{
	Follower string "json:\"follower\""
	Following string "json:\"following\""
}

type Photo struct{
	Id string "json:\"id\""
	User string "json\"userid\""
	Time time.Time "json\"time\""
}

type Like struct{
	User string "json\"userid\""
	Photo string "json\"photoid\""
}


type Comment struct{
	Id string "json\"id\""
	User string "json\"userid\""
	Photo string "json\"photoid\""
	Comment string "json\"comment\""
}

type Post struct{
	User string "json\"photo\""
	Photo string "json\"photo\""
	Likes string "json\"likes\""
	Comments string "json\"likes\""
	Time time.Time "json\"time\""
}

type AggregatedData struct {
    User     string      `json:"User"`
    Photo    string      `json:"Photo"`
    Likes    []string    `json:"Likes"`
    Comments []string    `json:"Comments"`
    Time     time.Time   `json:"Time"`
}