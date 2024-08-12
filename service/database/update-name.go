package database


func (db *appdbimpl) UpdateName(name string, id string) (error) {
	query := "UPDATE users SET username=? WHERE id=?"
	_, err := db.c.Exec(query,name, id)
	return err

}