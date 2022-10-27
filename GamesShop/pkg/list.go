package pkg

import "fmt"

func PrintGamesList(ls []*Game) {
	for i, v := range ls {
		fmt.Printf("%d. ", i+1)
		v.Print()
	}
}
