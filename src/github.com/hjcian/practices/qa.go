package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readQuestions(fpath string) (*[]string, *[]int) {
	csvfile, err := os.Open(fpath)
	if err != nil {
		fmt.Println("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	questions := make([]string, len(records))
	answers := make([]int, len(records))
	for idx := range records {
		questions[idx] = records[idx][0]
		ans, err := strconv.Atoi(records[idx][1])
		if err != nil {
			fmt.Println("answer sheet should be an integer.", err)
			os.Exit(1)
		}
		answers[idx] = ans
	}
	return &questions, &answers
}

func calcScore(userAns, examAns *[]int) int {
	score := 0
	for idx, user := range *userAns {
		if user == (*examAns)[idx] {
			score++
		}
	}
	return score
}

func readUserInput() int {
	stdinReader := bufio.NewReader(os.Stdin)
	ret := 0
	for {
		text, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("something went wrong... ", err)
		}
		text = strings.Replace(text, "\r\n", "", -1)
		ans, err := strconv.Atoi(text)
		if err == nil {
			ret = ans
			break
		} else {
			fmt.Println("your input isn't a number, please try again.")
		}
	}
	return ret
}

func quickQnA(signal chan<- bool, questions *[]string, userAns *[]int) {
	for idx, qs := range *questions {
		fmt.Printf("Q%v: %v = ", idx+1, qs)
		ans := readUserInput()
		(*userAns)[idx] = ans
	}
	signal <- true
}

func main() {
	// parse argument
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Println("error: one positional argument is required for parsing the question.")
		os.Exit(1)
	}
	// prepare questions, answers, user input container
	questions, answers := readQuestions(argsWithoutProg[0])
	fmt.Println(questions)
	fmt.Println(answers)
	userAns := make([]int, len(*questions))
	// ready go
	fmt.Println("You have 30 seconds to finish this exam, the exam will begin after 3 sec.")
	time.Sleep(3 * time.Second)
	fmt.Println("Start!")
	signal := make(chan bool, 1)
	go quickQnA(signal, questions, &userAns)
	select {
	case <-signal:
		fmt.Println("Congrats! You finish the exam on time!")
	case <-time.After(time.Second * 3):
		fmt.Println()
		fmt.Println("Time's up! So sad! you aren't complete the exam on time!")
	}
	// compute the score
	score := calcScore(&userAns, answers)
	fmt.Printf("Total score: %v/%v \n", score, len(*questions))
}
