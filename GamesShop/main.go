package main

import "awesomeProject/GamesShop/pkg"

func main() {
	gList := []*pkg.Game{
		{
			Title: "Minecraft",
			Price: 5,
		},
		{
			Title: "World of warcraft",
			Price: 19,
		},
		{
			Title: "Elite Dangerous",
			Price: 54,
		},
	}

	pkg.PrintGamesList(gList)
}
