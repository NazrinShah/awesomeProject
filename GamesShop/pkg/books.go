package pkg

import "fmt"

type Book struct {
	Title string
	Price int
}

func (b Book) Print() {
	fmt.Println(fmt.Sprintf("%s, $%d", b.Title, b.Price))
}
