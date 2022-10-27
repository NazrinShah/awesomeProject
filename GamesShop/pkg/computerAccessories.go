package pkg

import "fmt"

type ComputerAccessory struct {
	Title string
	Price int
}

func (c ComputerAccessory) Print() {
	fmt.Println(fmt.Sprintf("%s, $%d", c.Title, c.Price))
}
