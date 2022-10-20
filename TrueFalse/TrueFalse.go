package main

import (
	"fmt"
	"math/rand"
)

const (
	ENUM_QUIT = -1
	ENUM_HINT = -2
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

		if guess == ENUM_QUIT {
			fmt.Println("Thanks for playing!")
			break
		} else if guess == ENUM_HINT {
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
