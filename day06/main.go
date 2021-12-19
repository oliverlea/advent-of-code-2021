package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_FILE = "input"

const START_DAYS = 8
const RESTART_DAYS = 6

func simulate(ageCount map[int]int, days int) int {
	for day := 0; day < days; day++ {
		nextAgeCount := make(map[int]int)
		for age := 0; age <= START_DAYS; age++ {
			count := ageCount[age]
			if age == 0 {
				nextAgeCount[START_DAYS] += count
				nextAgeCount[RESTART_DAYS] += count
			} else {
				nextAgeCount[age-1] += count
			}
		}
		ageCount = nextAgeCount
	}
	sum := 0
	for _, count := range ageCount {
		sum += count
	}
	return sum
}

func main() {
	input := utils.ReadStrings(INPUT_FILE)[0]
	splitInput := strings.Split(input, ",")
	ages := make([]int, len(splitInput))
	for i, cur := range splitInput {
		ages[i], _ = strconv.Atoi(cur)
	}

	ageCount := make(map[int]int)
	for i := 0; i <= START_DAYS; i++ {
		ageCount[i] = 0
	}
	for _, age := range ages {
		ageCount[age]++
	}
	fmt.Println(simulate(ageCount, 80))
	fmt.Println(simulate(ageCount, 256))
}
