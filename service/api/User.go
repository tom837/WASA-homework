package API

type User struct {
	UserName string "json:\"name\""
	ID       string "json:\"ID\""
	Email    string "json:\"email\""
}
