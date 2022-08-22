package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret number is ", secretNumber)

	for {
		fmt.Println("Please input your guess")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Reading input occurs an error")
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input is invalid, please input data whose type is int")
			continue
		}
		//fmt.Println("your guess is", guess)

		if guess == secretNumber {
			fmt.Println("You are right, ", secretNumber, " is the corret answer")
			break
		} else if guess < secretNumber {
			fmt.Println(guess, " is lower than the correct answer")
		} else if guess > secretNumber {
			fmt.Println(guess, " is higher than the correct answer")
		}
	}

}
