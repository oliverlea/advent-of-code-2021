package main

import (
	"advent/utils"
	"fmt"
	"sort"
)

const INPUT_FILE = "input"

var INVALID_SCORE = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}
var TOKENS = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}
var COMPLETION_SCORE = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

type stack []rune

func (s stack) push(r rune) stack {
	return append(s, r)
}

func (s stack) pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) peek() (rune, bool) {
	if len(s) == 0 {
		return 0, false
	}
	return s[len(s)-1], true
}

func parse(instruction string) rune {
	var (
		s         = make(stack, 0, 100)
		expecting rune
	)
	for _, cur := range instruction {
		if right, isLeft := TOKENS[cur]; isLeft {
			s = s.push(right)
		} else {
			// Is right, is it the one we were expecting?
			s, expecting = s.pop()
			if cur != expecting { // Not adjacent in charset
				return cur
			}
		}
	}
	return 0
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func complete(instruction string) string {
	var (
		s = make(stack, 0, 100)
	)
	for _, cur := range instruction {
		if right, isLeft := TOKENS[cur]; isLeft {
			s = s.push(right)
		} else {
			// In this case we know the input is valid
			s, _ = s.pop()
		}
	}
	return reverse(string(s))
}

func part1(ss []string) {
	totalScore := 0
	for _, s := range ss {
		invalidRune := parse(s)
		curScore, exists := INVALID_SCORE[invalidRune]
		if exists {
			totalScore += curScore
		}
	}
	fmt.Println(totalScore)
}

func part2(ss []string) {
	scores := make([]int, 0, len(ss))
	for _, instruction := range ss {
		if parse(instruction) != 0 {
			continue
		}
		completedSuffix := complete(instruction)
		curScore := 0
		for _, r := range completedSuffix {
			curScore *= 5
			curScore += COMPLETION_SCORE[r]
		}
		scores = append(scores, curScore)
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

func main() {
	lines := utils.ReadStrings(INPUT_FILE)
	part1(lines)
	part2(lines)
}
