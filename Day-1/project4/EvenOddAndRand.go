package main

import (
	"fmt"
	"math/rand"
)

//	func EvenOdd(num int) {
//		if num%2 == 0 {
//			fmt.Println("Given number is Even")
//		} else {
//			fmt.Println("It is an odd")
//		}
//	}
func main() {
	//even&odd
	// num := 10
	// EvenOdd(num)

	// Create a simple guessing game.
	// . Generate a random number between 1 and 10, and let the user guess the number.
	// (use rand.Intn() method to generate random numbers)

	rand := rand.Intn(10)
	fmt.Println("Random number is", rand)
	var guessNumber int
	var oneMoreAttempt int

	for {
		fmt.Print("Enter number to guess")
		fmt.Scan(&guessNumber)
		oneMoreAttempt++
		if guessNumber == rand {
			fmt.Printf("guessed right %d \n", oneMoreAttempt)
			return
		} else if guessNumber < rand {
			fmt.Println("Too low")
		} else {
			fmt.Println("Too high")

		}
	}
}
