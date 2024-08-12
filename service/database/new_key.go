package database

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
)

func generateNewKey(db *appdbimpl, prefix string, database string) (string, error) {
	// Query to find the highest key
	var maxKey string
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", database)
	err := db.c.QueryRow(query).Scan(&maxKey)
	if err != nil {
		if err == sql.ErrNoRows {
			// No entries exist yet, start from u1
			return fmt.Sprintf("%s1", prefix), nil
		}
		// Handle other errors
		return "", err
	}

	// Extract the numeric part from the key
	re := regexp.MustCompile(fmt.Sprintf(`^%s(\d+)$`, prefix))
	matches := re.FindStringSubmatch(maxKey)
	if len(matches) != 2 {
		// The key format is not as expected
		return "", fmt.Errorf("unexpected key format: %s", maxKey)
	}

	// Convert the numeric part to an integer and increment it
	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return "", err
	}

	// Generate the new key
	newKey := fmt.Sprintf("%s%d", prefix, num+1)
	return newKey, nil
}
