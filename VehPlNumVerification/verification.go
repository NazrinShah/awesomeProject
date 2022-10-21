package main

import (
	"fmt"
	"strconv"
	"strings"
)

var charToIntMap = map[byte]int{
	'A': 1,
	'B': 2,
	'C': 3,
	'D': 4,
	'E': 5,
	'F': 6,
	'G': 7,
	'H': 8,
	'I': 9,
	'J': 10,
	'K': 11,
	'L': 12,
	'M': 13,
	'N': 14,
	'O': 15,
	'P': 16,
	'Q': 17,
	'R': 18,
	'S': 19,
	'T': 20,
	'U': 21,
	'V': 22,
	'W': 23,
	'X': 24,
	'Y': 25,
	'Z': 26,
}

var checkSumMap = map[byte]int{
	'A': 0,
	'Z': 1,
	'Y': 2,
	'X': 3,
	'U': 4,
	'T': 5,
	'S': 6,
	'R': 7,
	'P': 8,
	'M': 9,
	'L': 10,
	'K': 11,
	'J': 12,
	'H': 13,
	'G': 14,
	'E': 15,
	'D': 16,
	'C': 17,
	'B': 18,
}

func isLetter(s byte) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

func isNumber(s byte) bool {
	return s >= '0' && s <= '9'
}
func validateInput(input string) (bool, string, []byte, []int) {
	prefix := []byte{}
	nums := []int{}
	inLen := len(input)

	/*
		len min 3 max 8
		get letters for prefix
			- min 1 max 3
		get last letter, must be letter
	*/
	if inLen < 3 || inLen > 8 {
		return false, "Invalid plate number length", nil, nil
	}

	if !isLetter(input[inLen-1]) {
		return false, "Last char in plate number is NOT an alphabet", nil, nil
	}

	i := 0

	for ; i < 3; i++ {
		c := input[i]
		if !isLetter(c) {
			break
		}

		prefix = append(prefix, c)
	}

	for ; i < inLen-1; i++ {
		if v, err := strconv.Atoi(string(input[i])); err == nil {
			nums = append(nums, v)
			continue
		}

		return false, "Invalid plate number", nil, nil
	}

	return true, "", prefix, nums
}

func getInput() string {
	input := ""

	for {
		fmt.Printf("Please enter a plate number: ")
		_, err := fmt.Scan(&input)
		fmt.Println()

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		}

		input = strings.ToUpper(input)

		break
	}

	return input
}

func isValidPlateNumber(prefix []byte, nums []int, check byte) bool {
	i, j := 0, 0
	res := 0
	base := make([]int, 6)
	multiplier := []int{9, 4, 5, 4, 3, 2}
	pLen := len(prefix)

	if pLen == 3 {
		i = 1
	} else if pLen == 1 {
		j = 1
	}

	for ; i < pLen; i++ {
		base[j] = charToIntMap[prefix[i]]
	}

	j = 5
	for k := len(nums) - 1; k >= 0; k-- {
		base[j] = nums[k]
		j--
	}

	for i, v := range multiplier {
		res += base[i] * v
	}

	checksum := res % 19

	fmt.Println(fmt.Sprintf("calculated checksum is %d, expecting %d", checksum, checkSumMap[check]))
	return checkSumMap[check] == checksum
}

func main() {

	input, prefix, nums := "", []byte{}, []int{}

	for {
		input = getInput()
		valid, errMsg, _prefix, _nums := validateInput(input)

		if !valid {
			fmt.Println("Error processing input: ", errMsg)
			continue
		}

		prefix = _prefix
		nums = _nums

		break
	}

	if isValidPlateNumber(prefix, nums, input[len(input)-1]) {
		fmt.Println("The vehicle plate given is correct!")
	} else {
		fmt.Println("The vehicle plate given is not correct!")
	}
}
