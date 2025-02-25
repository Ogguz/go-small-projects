package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 7
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:] // 1 because it starts with program name
	if len(args) != 1 {
        fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Please insert a number!")
		return
	}

	if guess < 0 {
		fmt.Println("Please insert a positive number!")
		return
	}

	for turn := 0; turn < maxTurns ; turn ++ {
		n := rand.Intn(guess+1)
		fmt.Printf("Picked: %d \n",n)
		if n == guess {
			fmt.Println("YOU WIN!")
			return
		}
		fmt.Printf("Again!\n")
	}

	fmt.Println("Game is over...")

}
