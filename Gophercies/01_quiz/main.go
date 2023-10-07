package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the quiz wizzard!")

	timer := time.Duration(30)

	f, err := os.Open("problems.csv")
	ErrorCheck(err)

	defer f.Close()

	csvReader := csv.NewReader(f)

	var score = 0
	var ans = ""

	record, err := csvReader.ReadAll()

	for _, rec := range record {
		// Assuming you want to access the first column (index 0)

		if timer.Seconds() == 30 {
			break
		}

		columnValue := rec[0]

		fmt.Printf("What %s , sir ?\n", columnValue)
		fmt.Print("Enter the answer: ")
		fmt.Scan(&ans)

		if ans == string(rec[1]) {
			score++
		}

	}

	fmt.Println("Your score is: ", score)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
