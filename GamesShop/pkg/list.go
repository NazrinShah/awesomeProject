package pkg

import "fmt"

/*
Item interface
*/
type Item interface {
	Print()
}

type List []Item //!< list of items

/*
Prints items in list
*/
func (l *List) Print() {
	for i, v := range *l {
		fmt.Printf("%d. ", i+1)
		v.Print()
	}
}
