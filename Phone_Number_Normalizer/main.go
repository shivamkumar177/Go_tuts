package main

import (
	"database/sql"
	"fmt"
	"regexp"
)

// implement the fucntion and write test for the fucntions
func NormalizePhoneNumbers(givenNumber []string) []string {
	normalizedNumber := make([]string, len(givenNumber))
	for _, number := range givenNumber {
		re := regexp.MustCompile("\\D")
		normalizedNumber = append(normalizedNumber, re.ReplaceAllString(number, ""))
	}
	return normalizedNumber
}

func createTable(db *sql.DB) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS playing_with_neon(id SERIAL PRIMARY KEY, name TEXT NOT NULL, value REAL);")
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func insertRec(db *sql.DB) (int64, error) {
	query := "INSERT INTO playing_with_neon(name, value) values ('shivam', 23)"
	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	lastInserted, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInserted, nil
}

func main() {
	givenNumber := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 782",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	fmt.Println(NormalizePhoneNumbers(givenNumber))

	db := DbConn()
	insertRec(db)
}
