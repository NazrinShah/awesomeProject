package pkg

import "fmt"

/*
Book struct to store book details
*/
type Book struct {
	Title string
	Price int
}

/*
Prints book details
*/
func (b *Book) Print() {
	fmt.Println(fmt.Sprintf("%s, $%d", b.Title, b.Price))
}
