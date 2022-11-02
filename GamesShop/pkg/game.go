package pkg

import "fmt"

/*
Game struct used to store Game details
*/
type Game struct {
	Title string
	Price float64
}

/*
Prints Game details
*/
func (g *Game) Print() {
	fmt.Println(fmt.Sprintf("%s, $%.2f", g.Title, g.Price))
}
