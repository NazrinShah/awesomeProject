package main

import (
	"errors"
	"fmt"
)

func getInput() int {
	var num int

	for {

		fmt.Printf("Enter an integer between 0-9: ")
		_, err := fmt.Scan(&num)
		fmt.Println()

		if err != nil {
			fmt.Printf("Error processing input: %v\n", err)
			continue
		}

		if num < 0 || num > 9 {
			fmt.Printf("Invalid input. Expecting number between 0-9 but got %d\n", num)
			continue
		}

		break
	}

	return num
}

func add(lhs, rhs int) (int, error) {
	return lhs + rhs, nil
}

func subtract(lhs, rhs int) (int, error) {
	return lhs - rhs, nil
}

func multiply(lhs, rhs int) (int, error) {
	return lhs * rhs, nil
}

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return -1, errors.New("denominator cannot be 0")
	}
	return lhs / rhs, nil
}

func getOp() int {
	op := OP_ADD

	for {
		fmt.Printf("1: addition\n2: subtraction\n3: multiplication\n4: division\nEnter an integer representing the operation: ")
		_, err := fmt.Scan(&op)
		fmt.Println()

		if err != nil {
			fmt.Println("Error processing input: %v", err)
			continue
		}

		if op < 1 || op > 4 {
			fmt.Printf("Invalid input. Expecting input between 1-4 but got %d\n", op)
			continue
		}

		break
	}

	return op
}

const (
	OP_ADD      = 1
	OP_SUBTRACT = 2
	OP_MULTIPLY = 3
	OP_DIVIDE   = 4
)

func main() {
	ops := map[int]func(int, int) (int, error){
		OP_ADD:      add,
		OP_SUBTRACT: subtract,
		OP_MULTIPLY: multiply,
		OP_DIVIDE:   divide,
	}

	opSymbol := map[int]string{
		OP_ADD:      "+",
		OP_SUBTRACT: "-",
		OP_MULTIPLY: "*",
		OP_DIVIDE:   "/",
	}

	inputs := make([]int, 2)
	inputs[0] = getInput()
	op := getOp()
	inputs[1] = getInput()

	res, err := ops[op](inputs[0], inputs[1])
	if err != nil {
		fmt.Println("Cannot perform calculation: ", err)
	} else {
		fmt.Println(fmt.Sprintf("%d %s %d = %v", inputs[0], opSymbol[op], inputs[1], res))
	}
}
