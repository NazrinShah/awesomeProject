package main

import (
	"fmt"
	"math/rand"
)

func main() {
	actual := int(rand.Intn(99) + 1)
	var guess int

	for {
		fmt.Println("Enter an integer value (-1 to quit): ")
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Printf("Error processing input: %v\n", err)
			continue
		}

		if guess == -1 {
			fmt.Println("Thanks for playing!")
			break
		} else if guess == -2 {
			fmt.Printf("Answer is: %d\n", actual)
		}

		if guess == actual {
			fmt.Println("Well Done! Your guess is correct")
			break
		} else if guess < actual {
			fmt.Println("Too low, try again next time!")
		} else {
			fmt.Println("Too high, try again next time!")
		}
	}
}
