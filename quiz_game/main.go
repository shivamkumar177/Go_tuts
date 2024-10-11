package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func parseCsvRecords(records [][]string) []problem {
	problemList := make([]problem, len(records))
	for i, line := range records {
		problemList[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return problemList
}
func main() {
	csvFile := flag.String("csv", "problems.csv", "csv file of question,answer file")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("files data", records)

	problems := parseCsvRecords(records)
	counter := 0
	for _, eachProblem := range problems {
		fmt.Println("Question: ", eachProblem.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == eachProblem.a {
			fmt.Println("correct")
			counter += 1
		}
	}
	fmt.Printf("You got %d from %d\n", counter, len(problems))
}
