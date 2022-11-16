package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func write(fileName string, res chan string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		res <- scanner.Text()
	}

	file.Close()
}

func main() {
	resChan1 := make(chan string)
	resChan2 := make(chan string)

	res1 := make([]string, 0)
	res2 := make([]string, 0)

	go write("file1.txt", resChan1)
	go write("file2.txt", resChan2)
	line := ""
	shouldBreak := false

	for !shouldBreak {
		select {
		case line = <-resChan1:
			if line != "" {
				res1 = append(res1, line)
			}
		case line = <-resChan2:
			res2 = append(res2, line)
		case <-time.After(1 * time.Millisecond):
			fmt.Println("Operation has timed out")
			shouldBreak = true
		}
	}

	fmt.Println("File1: ", res1)
	fmt.Println("File2: ", res2)
}
