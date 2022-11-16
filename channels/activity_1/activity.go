package main

import "fmt"

func hello(num int, done chan bool) {
	fmt.Println("Hello world goroutine ", num)
	done <- true
}

func main() {
	done := make(chan bool)
	n := 4

	for i := 0; i < n; i++ {
		go hello(i+1, done)
	}

	for n > 0 {
		select {
		case <-done:
			n--
		}
	}
}
