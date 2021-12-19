package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadLineCh(fn string, ch chan string, done chan bool) {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal("Could not open file:", fn)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- string(scanner.Text())
	}
	done <- true
}

func ReadInts(fn string) []int {
	inCh, doneCh := make(chan string, 100), make(chan bool)
	go ReadLineCh(fn, inCh, doneCh)

	var output []int
loop:
	for {
		select {
		case line := <-inCh:
			cur, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			output = append(output, cur)
		case <-doneCh:
			close(inCh)
			close(doneCh)
			break loop
		}
	}
	return output
}
