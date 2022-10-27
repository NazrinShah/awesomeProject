package main

import (
	"fmt"
)

func getInput() int {
	in := 0

	for {
		fmt.Printf("Please enter a non-negative number: ")
		_, err := fmt.Scan(&in)
		fmt.Println()

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		} else if in < 0 {
			continue
		}

		break
	}

	return in
}

func factorial(n int) int {
	if n < 0 {
		return 1
	}

	res := 1

	for n > 0 {
		res *= n
		n--
	}

	return res
}

func main() {
	fmt.Println(factorial(getInput()))
}
