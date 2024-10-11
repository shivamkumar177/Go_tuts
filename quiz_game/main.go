package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
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
	timeLimit := flag.Int("time", 30, "time limit for the quiz")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	counter := 0
	for _, eachProblem := range problems {
		fmt.Println("Question: ", eachProblem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You got %d from %d\n", counter, len(problems))
			return
		case answer := <-answerCh:
			if answer == eachProblem.a {
				fmt.Println("correct")
				counter += 1
			}
		}
	}
	fmt.Printf("You got %d from %d\n", counter, len(problems))
}
