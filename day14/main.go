package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strings"
)

const INPUT_FILE = "testinput"

type pair struct {
	first  rune
	second rune
}

type state struct {
	chain []rune
	rules map[pair]rune
}

func parseInput(lines []string) state {
	var (
		initialChain = []rune(lines[0])
		rules        = make(map[pair]rune)
	)
	lines = lines[2:]
	for _, l := range lines {
		var (
			splitL  = strings.Split(l, " -> ")
			first   = rune(splitL[0][0])
			second  = rune(splitL[0][1])
			product = rune(splitL[1][0])
		)
		rules[pair{first, second}] = product
	}
	return state{initialChain, rules}
}

func (s *state) step() {
	var (
		newChain = make([]rune, len(s.chain)*2)
		curLoc   = 0
	)
	for i := 0; i < len(s.chain)-1; i++ {
		newChain[curLoc] = s.chain[i]
		curLoc++

		p := pair{s.chain[i], s.chain[i+1]}
		product, exists := s.rules[p]
		if exists {
			newChain[curLoc] = product
		}
		curLoc++
	}
	newChain[curLoc] = s.chain[len(s.chain)-1]
	curLoc++
	newChain = newChain[:curLoc]
	s.chain = newChain
}

func part1(s state) {
	for i := 0; i < 10; i++ {
		(&s).step()
	}
	counts := make(map[rune]int)
	for _, r := range s.chain {
		counts[r]++
	}
	var (
		min, max = math.MaxInt, 0
		// minR, maxR rune
	)
	for _, count := range counts {
		if count < min {
			min = count
			// minR = r
		}
		if count > max {
			max = count
			// maxR = r
		}
	}
	fmt.Println(min, max)
	fmt.Println(max - min)
}

func part2(s state) {
	// Our target is to expand every possible expansion for half the required steps 40 -> 20.
	// From this, then simulate the 20 steps of the main polymer string, and count
	// the expansions as proxy pointers into our cached lookup of expansions over the
	// next 20 steps :)
	simulated := make(map[pair][]rune)
	simulationCounts := make(map[pair]map[rune]int)
	for pair := range s.rules {
		simulateState := state{[]rune{pair.first, pair.second}, s.rules}
		for i := 0; i < 20; i++ {
			(&simulateState).step()
		}
		simulated[pair] = simulateState.chain
		simulationCounts[pair] = make(map[rune]int)
		for _, r := range simulateState.chain {
			simulationCounts[pair][r]++
		}
	}

	for i := 0; i < 20; i++ {
		(&s).step()
	}

	totalCounts := make(map[rune]int)
	for i := 0; i < len(s.chain)-1; i++ {
		p := pair{s.chain[i], s.chain[i+1]}
		for r, count := range simulationCounts[p] {
			totalCounts[r] += count
		}
	}

	var (
		min, max = math.MaxInt, 0
		// minR, maxR rune
	)
	for _, count := range totalCounts {
		if count < min {
			min = count
			// minR = r
		}
		if count > max {
			max = count
			// maxR = r
		}
	}
	fmt.Println(min, max)
	fmt.Println(max - min)
}

func main() {
	lines := utils.ReadStrings(INPUT_FILE)
	state := parseInput(lines)
	part1(state)

	lines = utils.ReadStrings(INPUT_FILE)
	state = parseInput(lines)
	part2(state)
}
