package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

func ReadStrings(fn string) []string {
	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(bs), "\n")
	trimmedLines := make([]string, len(lines))
	for i, line := range lines {
		trimmedLines[i] = strings.TrimSpace(line)
	}
	return trimmedLines
}

func Limit(xs []int) (int, int) {
	var min, max = math.MaxInt, 0
	for _, x := range xs {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max
}
