package main

import (
	"flag"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "csv file of question,answer file")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		panic(err)
	}
	_ = file
}
