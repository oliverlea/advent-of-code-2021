package main

import (
	"advent/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const INPUT_FILE = "input"

func median(xs []int) int {
	midIndex := len(xs) / 2
	if midIndex%2 == 0 {
		return xs[midIndex]
	}
	return (xs[midIndex-1] + xs[midIndex]) / 2
}

func movementRequiredLinear(xs []int, position int) int {
	moves := 0
	for _, val := range xs {
		diff := val - position
		if diff < 0 {
			diff *= -1
		}
		moves += diff
	}
	return moves
}

func movementRequiredStep(xs []int, position int) int {
	moves := 0
	for _, val := range xs {
		diff := val - position
		if diff < 0 {
			diff *= -1
		}
		moves += (diff * (diff + 1)) / 2
	}
	return moves
}

func part1(xs []int) {
	targetPosition := median(xs)
	moves := movementRequiredLinear(xs, targetPosition)
	fmt.Println(moves)
}

func part2(xs []int) {
	min, max := xs[0], xs[len(xs)-1]
	best := math.MaxInt
	for i := min; i <= max; i++ {
		moves := movementRequiredStep(xs, i)
		if moves < best {
			best = moves
		}
	}
	fmt.Println(best)
}

func main() {
	input := strings.Split((utils.ReadStrings(INPUT_FILE)[0]), ",")
	numbers := make([]int, len(input))
	for i, valStr := range input {
		numbers[i], _ = strconv.Atoi(valStr)
	}

	sort.Ints(numbers)
	part1(numbers)
	part2(numbers)
}
