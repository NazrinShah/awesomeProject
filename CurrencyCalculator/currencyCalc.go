package main

import (
	"fmt"
)

const (
	CoinTypeOneDollar = iota
	CoinTypeFiftyCent
	CoinTypeTwentyCent
	CoinTypeTenCent
	CoinTypeFiveCent
)

var baseMap = map[int]float64{
	CoinTypeOneDollar:  1.0,
	CoinTypeFiftyCent:  0.5,
	CoinTypeTwentyCent: 0.2,
	CoinTypeTenCent:    0.1,
	CoinTypeFiveCent:   0.05,
}

func getCoinNum(coinType int) int {
	var num int
	coinStr := ""

	switch coinType {
	case CoinTypeOneDollar:
		coinStr = "1-dollar"
		break
	case CoinTypeFiftyCent:
		coinStr = "50-cent"
		break
	case CoinTypeTwentyCent:
		coinStr = "20-cent"
		break
	case CoinTypeTenCent:
		coinStr = "10-cent"
		break
	default:
		coinStr = "5-cent"
	}

	for {
		fmt.Printf("Enter number of %s coins (must be at least 0): ", coinStr)
		_, err := fmt.Scan(&num)

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		} else if num < 0 {
			fmt.Println("Must have at least 0 coins")
			continue
		}

		break
	}

	return num
}

func main() {
	coinTypes := []int{CoinTypeOneDollar, CoinTypeFiftyCent, CoinTypeTwentyCent, CoinTypeTenCent, CoinTypeFiveCent}
	total := float64(0)
	for _, v := range coinTypes {
		total += float64(getCoinNum(v)) * baseMap[v]
	}

	fmt.Println("Total: ", total)

	notes := int64(total) / 2
	fmt.Println(fmt.Sprintf("Can be converted into %d 2-dollar notes with %.2f remaining", notes, total-float64(notes*2)))
}
