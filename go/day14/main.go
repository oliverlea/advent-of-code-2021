package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strings"
)

const INPUT_FILE = "input"

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

func limitCounts(m map[rune]int) (int, int) {
	min, max := math.MaxInt, 0
	for _, count := range m {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	return min, max
}

func part1(s state) {
	for i := 0; i < 10; i++ {
		(&s).step()
	}
	counts := make(map[rune]int)
	for _, r := range s.chain {
		counts[r]++
	}

	min, max := limitCounts(counts)
	fmt.Println(min, max)
	fmt.Println(max - min)
}

func part2(s state) {
	// Expand every possible expansion for half the required steps 40 -> 20.
	// From this, then simulate the 20 steps of the main polymer string, and count
	// the expansions as proxy pointers into our cached lookup of expansions over the
	// next 20 steps
	simulationCounts := make(map[pair]map[rune]int)
	for pair := range s.rules {
		simulatedState := state{[]rune{pair.first, pair.second}, s.rules}
		for i := 0; i < 20; i++ {
			(&simulatedState).step()
		}
		trimmedChain := simulatedState.chain
		if len(trimmedChain) > 2 {
			// Chop off the end so the chain is just the expanded section
			trimmedChain = trimmedChain[1 : len(trimmedChain)-1]
		} else {
			trimmedChain = []rune{}
		}
		simulationCounts[pair] = make(map[rune]int)
		for _, r := range trimmedChain {
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
		totalCounts[s.chain[i]]++
	}
	totalCounts[s.chain[len(s.chain)-1]]++

	min, max := limitCounts(totalCounts)
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
