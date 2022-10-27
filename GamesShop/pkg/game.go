package pkg

import "fmt"

type Game struct {
	Title string
	Price float64
	_     struct{}
}

func (g *Game) Print() {
	fmt.Println(fmt.Sprintf("%s, $%.2f", g.Title, g.Price))
}
