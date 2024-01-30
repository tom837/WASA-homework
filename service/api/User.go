package API

import (
	"fmt"
)


type User struct{
	UserName string 'json:"name"'
	ID string 'json:"ID"'
	Email string 'json:"email"'
}

