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
	secretNum := rand.Intn(maxNum)
	//fmt.Println("the secret number is", secretNum)

	fmt.Println("Guess the number between 1 and 100")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Your guess is", guess)

		if guess > secretNum {
			fmt.Println("Your guess is too high")
		} else if guess < secretNum {
			fmt.Println("Your guess is too low")
		} else {
			fmt.Println("Congratulation! You got the number")
			break
		}
	}
}
