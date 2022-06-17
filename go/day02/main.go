package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_FILE = "input"

type instruction struct {
	direction string
	amount    uint64
}

func parseInput(s string) instruction {
	splitInstruction := strings.Split(s, " ")
	amount, _ := strconv.ParseUint(splitInstruction[1], 10, 0)
	return instruction{splitInstruction[0], amount}
}

func part1(instructions []instruction) {
	var (
		hor   uint64 = 0
		depth uint64 = 0
	)
	for _, cur := range instructions {
		switch cur.direction {
		case "forward":
			hor += cur.amount
		case "down":
			depth += cur.amount
		case "up":
			depth -= cur.amount
		}
	}
	fmt.Println(hor * depth)
}

func part2(instructions []instruction) {
	var (
		hor   uint64 = 0
		depth uint64 = 0
		aim   uint64 = 0
	)
	for _, cur := range instructions {
		switch cur.direction {
		case "forward":
			hor += cur.amount
			depth += aim * cur.amount
		case "down":
			aim += cur.amount
		case "up":
			aim -= cur.amount
		}
	}
	fmt.Println(hor * depth)

}

func main() {
	input := utils.ReadStrings(INPUT_FILE)
	instructions := make([]instruction, len(input))
	for i, cur := range input {
		instructions[i] = parseInput(cur)
	}

	part1(instructions)
	part2(instructions)
}
