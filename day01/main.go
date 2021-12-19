package main

import (
	"advent/utils"
	"fmt"
	"math"
)

const INPUT_FILE = "input"

func part1(values []int) {
	prev, increases := math.MaxInt, 0
	for _, cur := range values {
		fmt.Println(cur)
		if cur > prev {
			increases++
		}
		prev = cur
	}
	fmt.Println(increases)
}

func part2(values []int) {
	increases, prevSum := 0, math.MaxInt
	for i := 0; i < len(values)-2; i++ {
		curSum := values[i] + values[i+1] + values[i+2]

		if curSum > prevSum {
			increases++
		}
		prevSum = curSum
	}
	fmt.Println(increases)
}

func main() {
	values := utils.ReadInts(INPUT_FILE)

	part1(values)
	part2(values)
}
