package main

import (
	"fmt"
	"regexp"
)

// implement the fucntion and write test for the fucntions
func normalizePhoneNumbers(givenNumber []string) []string {
	normalizedNumber := make([]string, len(givenNumber))
	for _, number := range givenNumber {
		re := regexp.MustCompile("\\D")
		normalizedNumber = append(normalizedNumber, re.ReplaceAllString(number, ""))
	}
	return normalizedNumber
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
	fmt.Println(normalizePhoneNumbers(givenNumber))
}
