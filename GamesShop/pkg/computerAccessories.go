package pkg

import "fmt"

/*
Computer Accessory struct used to store accessory details
*/
type ComputerAccessory struct {
	Title string
	Price int
}

/*
Prints accessory details
*/
func (c *ComputerAccessory) Print() {
	fmt.Println(fmt.Sprintf("%s, $%d", c.Title, c.Price))
}
