package main

import "awesomeProject/GamesShop/pkg"

func main() {
	gList := []pkg.Item{
		pkg.Game{
			Title: "Minecraft",
			Price: 5,
		},
		pkg.Game{
			Title: "World of warcraft",
			Price: 19,
		},
		pkg.Game{
			Title: "Elite Dangerous",
			Price: 54,
		},
		pkg.Book{
			Title: "Candle in the tomb",
			Price: 20,
		},
		pkg.Book{
			Title: "Barney and Friends",
			Price: 10,
		},
		pkg.ComputerAccessory{
			Title: "Razer BT earpiece",
			Price: 159,
		},
		pkg.ComputerAccessory{
			Title: "Razer keyboard",
			Price: 110,
		},
		pkg.ComputerAccessory{
			Title: "Logitech Mouse",
			Price: 80,
		},
	}
	pkg.PrintGamesList(gList)
}
