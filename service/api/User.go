package api

import "time"

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
