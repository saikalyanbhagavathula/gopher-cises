package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"
)

var quizFileName, answer, anyKey string
var correctAnswers, totalQuestions uint16
var countDown time.Duration

func main() {
	fmt.Println("Lets Start The Quiz....!")
	fmt.Print("Enter the coundown for quiz: ")
	fmt.Scanln(&countDown)
	if countDown == 0 {
		countDown = 30
	}
	//fmt.Println("countdown: ", countDown)
	fmt.Print("Press ANy Key To Start A Quiz: ")
	fmt.Scanln(&anyKey)
	timer := time.NewTimer(countDown * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Time Completed")
				fmt.Println("Score was ", correctAnswers, "/", totalQuestions)
				os.Exit(1)
			}
		}
	}()
	err := startQuiz()
	if err != nil {
		fmt.Print(err)
	}

}

func startQuiz() (err error) {

	fmt.Print("Enter a file name: ")
	fmt.Scanln(&quizFileName)
	data, err := os.Open(quizFileName)
	defer data.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot Process further: ", err, quizFileName))
	}
	records, err := csv.NewReader(data).ReadAll()
	for _, statement := range records {
		fmt.Printf("what is " + statement[0] + " sir?, ")
		fmt.Scan(&answer)
		if answer == statement[1] {
			fmt.Println("Congratulations the answer was correct")
			correctAnswers += 1
		}
		totalQuestions += 1
	}
	fmt.Println("Score was ", correctAnswers, "/", totalQuestions)
	return
}
