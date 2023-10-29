package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the quiz wizzard!")

	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")

	f, err := os.Open("problems.csv")
	ErrorCheck(err)

	defer f.Close()

	csvReader := csv.NewReader(f)

	var score = 0

	record, err := csvReader.ReadAll()
	ErrorCheck(err)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
Loop:
	for _, rec := range record {
		// Assuming you want to access the first column (index 0)
		columnValue := rec[0]
		fmt.Printf("What %s , sir ?\n", columnValue)
		ansch := make(chan string)
		go func() {
			var ans string
			fmt.Print("Enter the answer: ")
			fmt.Scan(&ans)
			ansch <- ans
		}()
		select {
		case <-timer.C:
			fmt.Println("Your score is: ", score)
			break Loop
		case ans := <-ansch:
			if ans == string(rec[1]) {
				score++
			}

		}
	}

	fmt.Println("Your score is: ", score)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
