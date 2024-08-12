package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetName(name string) (string, error) {
	newKey, err := generateNewKey(db, "u", "users")
	if err != nil {
		return "", err
	}
	query := "INSERT INTO users (id,username) VALUES (?,?)"
	_, err = db.c.Exec(query, newKey, name)
	if err != nil {
		return "", err
	}

	return newKey, err
}
