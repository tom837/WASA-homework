package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetName(id string) (string, error) {
	var name string
	query := "SELECT username FROM users WHERE id=?"
	err := db.c.QueryRow(query, id).Scan(&name)
	return name, err
}
