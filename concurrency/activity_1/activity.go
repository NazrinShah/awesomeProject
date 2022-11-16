package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Game struct {
	title string
	price float64
}

type GameShop struct {
	games     []*Game
	mtx       sync.Mutex
	inputLock sync.Mutex
	wg        sync.WaitGroup
}

const (
	DISPLAY = iota + 1
	ADD
	QUIT
)

var Shop = InitGameShop()

func InitGameShop() *GameShop {
	return &GameShop{
		games: []*Game{},
		mtx:   sync.Mutex{},
		wg:    sync.WaitGroup{},
	}
}

func (gs *GameShop) Add(g *Game) {
	gs.mtx.Lock()

	if gs.games == nil {
		gs.games = make([]*Game, 0)
	}

	gs.games = append(gs.games, g)

	gs.mtx.Unlock()
}

func (gs *GameShop) Display() {
	gs.mtx.Lock()

	if gs.games == nil {
		gs.games = make([]*Game, 0)
	}

	fmt.Println("Games List")
	fmt.Println("===================================")

	for i, g := range gs.games {
		fmt.Println(fmt.Sprintf("%d. %s: $%.2f", i+1, g.title, g.price))
	}

	fmt.Println("===================================")

	gs.mtx.Unlock()
}

func manageGameItems(id int) {
	defer Shop.wg.Done()
	in := 0
	locked := false

	for in != QUIT {

		if !locked {
			locked = Shop.inputLock.TryLock()

			if !locked {
				continue
			}
		}

		fmt.Println("Select action:\n1: Display Games \n2: Add Game\n3: Quit")
		_, err := fmt.Scan(&in)

		if err != nil {
			fmt.Println(fmt.Sprintf("[Thread %d] Error processing input: %+v", id, err))
			continue
		} else if in < DISPLAY || in > QUIT {
			fmt.Println(fmt.Sprintf("[Thread %d]Invalid input: Expecting value between %d and %d but received %d", id, DISPLAY, QUIT, in))
			continue
		}

		switch in {
		case DISPLAY:
			fmt.Println(in)
			Shop.Display()
		case ADD:
			game := &Game{}
			for {
				fmt.Printf("[Thread %d] Enter name of game: ", id)
				reader := bufio.NewReader(os.Stdin)
				game.title, err = reader.ReadString('\n')
				game.title = game.title[0 : len(game.title)-1]
				fmt.Println()

				if err != nil {
					fmt.Println(fmt.Sprintf("[Thread %d] Error processing input: %+v", id, err))
					continue
				} else if game.title == "" {
					fmt.Println(fmt.Sprintf("[Thread %d] Invalid input: empty string not allowed", id))
					continue
				}

				break
			}

			for {
				fmt.Printf("[Thread %d] Enter price of game: ", id)
				_, err := fmt.Scan(&game.price)
				fmt.Println()

				if err != nil {
					fmt.Println(fmt.Sprintf("[Thread %d] Error processing input: %+v", id, err))
				}

				break
			}

			Shop.Add(game)
		}

		Shop.inputLock.Unlock()
		locked = false
	}
}

func bot(done chan bool) {
	for {
		select {
		case <-done:
			break
		default:
			time.Sleep(5 * time.Second)
			Shop.Display()
		}
	}
}

func main() {
	n := 2
	for i := 0; i < n; i++ {
		Shop.wg.Add(1)
		go manageGameItems(i)
	}

	done := make(chan bool)
	go bot(done)
	Shop.wg.Wait()

	done <- true
}
