package pkg

import "fmt"

type Item interface {
	Print()
}

func PrintGamesList(ls []Item) {
	for i, v := range ls {
		fmt.Printf("%d. ", i+1)
		v.Print()
	}
}
