package main

import "fmt"

func squares(n int, res chan int) {
	sq := 0

	for n > 0 {
		i := n % 10

		sq += i * i

		n = n / 10
	}

	res <- sq
}

func cubes(n int, res chan int) {
	cb := 0

	for n > 0 {
		i := n % 10

		cb += i * i * i

		n = n / 10
	}

	res <- cb
}

func main() {
	in := 0
	res := make(chan int)

	for {
		fmt.Printf("Enter an integer value: ")
		_, err := fmt.Scan(&in)

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		}

		break
	}

	go squares(in, res)
	go cubes(in, res)

	fmt.Println(fmt.Sprintf("output = %d", <-res+<-res))
}
