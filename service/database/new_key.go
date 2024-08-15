package database

import (
	"database/sql"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func generateNewKey(db *appdbimpl, prefix string, database string) (string, error) {
	// Query to find the highest key
	query:= fmt.Sprintf("SELECT id FROM %s", database)
	rows, err:= db.c.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			// No entries exist yet, start from u1
			return fmt.Sprintf("%s1", prefix), nil
		}
		// Handle other errors
		return "", err
	}
	lst:= []int{}
	var id string
	for rows.Next() { //loop through all the users
		err = rows.Scan(&id)
		if err != nil {
			return "", err
		}
		// Extract the numeric part from the key
		re := regexp.MustCompile(fmt.Sprintf(`^%s(\d+)$`, prefix))
		matches := re.FindStringSubmatch(id)
		if len(matches) != 2 {
			// The key format is not as expected
			return "", fmt.Errorf("unexpected key format: %s", id)
		}
		// Convert the numeric part to an integer
		num, err := strconv.Atoi(matches[1])
		if err != nil {
			return "", err
		}
		lst = append(lst, num) //add the json to the final list
	}
	// Generate the new key
	max:=slices.Max(lst)
	newKey := fmt.Sprintf("%s%d", prefix, max+1)
	return newKey, nil
}
