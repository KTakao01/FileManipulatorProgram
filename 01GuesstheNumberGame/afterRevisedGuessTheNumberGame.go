package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var max, min int

	fmt.Println("Please enter the maximum and minimum values of the random number (e.g., 'm n').")
	fmt.Print("Enter max and min: ")
	_, err := fmt.Scanf("%d %d", &max, &min)
	if err != nil || max <= min {
		fmt.Println("Invalid input. Ensure number and max > min.")
		return
	}

	answer := generateRandomNumber(max, min)
	attempts := (max - min + 1) / 2
	score := playGame(answer, attempts)

	fmt.Printf("You got a score of: %d %%\n", score)
}

func generateRandomNumber(max, min int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(max-min+1) + min
}

func playGame(answer, attempts int) int {
	var guess int
	for i := 0; i < attempts; i++ {
		fmt.Print("Guess the number: ")
		fmt.Scanf("%d", &guess)

		if guess == answer {
			fmt.Println("Correct!")
			return (attempts - i) * 100 / attempts
		} else if guess > answer {
			fmt.Println("Incorrect! Your guess is too high.")
		} else {
			fmt.Println("Incorrect! Your guess is too low.")
		}
	}
	return 0
}
