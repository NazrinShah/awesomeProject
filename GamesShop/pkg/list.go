package pkg

import "fmt"

type Item interface {
	Print()
}

type List []Item

func (l *List) Print() {
	for i, v := range *l {
		fmt.Printf("%d. ", i+1)
		v.Print()
	}
}
