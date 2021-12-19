package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const INPUT_FILE = "input"
const BOARD_SIZE = 5

type score int
type board [][]int

type index struct {
	boardIndex int
	row        int
	col        int
}

func newBoard() board {
	b := make([][]int, BOARD_SIZE)
	for i := 0; i < BOARD_SIZE; i++ {
		b[i] = make([]int, BOARD_SIZE)
	}
	return b
}

func readInput(fn string) ([]board, []int) {
	lines := utils.ReadStrings(fn)
	drawsStrings := strings.Split(lines[0], ",")

	draws := make([]int, len(drawsStrings))
	for i, s := range drawsStrings {
		draws[i], _ = strconv.Atoi(s)
	}

	lines = lines[1:]
	boards := make([]board, 0)
	curBoard := newBoard()
	curRow := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			if curRow > 0 {
				// This board was actually used
				boards = append(boards, curBoard)
			}
			curBoard = newBoard()
			curRow = 0
			continue
		}
		words := strings.Split(line, " ")

		numbers := make([]int, BOARD_SIZE)
		curIndex := 0
		for _, w := range words {
			w = strings.ReplaceAll(w, " ", "")
			if len(w) > 0 {
				numbers[curIndex], _ = strconv.Atoi(w)
				curIndex++
			}
		}
		for x, cur := range numbers {
			curBoard[curRow][x] = cur
		}
		curRow++
	}
	if curRow > 0 {
		// This board was actually used
		boards = append(boards, curBoard)
	}

	return boards, draws
}

func rowSet(b board, index int) bool {
	for _, val := range b[index] {
		if val >= 0 {
			return false
		}
	}
	return true
}

func colSet(b board, index int) bool {
	for _, row := range b {
		if row[index] >= 0 {
			return false
		}
	}
	return true
}

func (b board) scoreBoard(draws []int) (score, int) {
	curDraw, curTurn := -1, -1
	completed := false
search:
	for turn, d := range draws {
		curDraw, curTurn = d, turn
		for y, row := range b {
			for x, val := range row {
				if val == d {
					b[y][x] = (val + 1) * -1
					if rowSet(b, y) || colSet(b, x) {
						completed = true
						break search
					}
				}
			}
		}
	}
	if !completed {
		return 0, -1
	}
	sum := 0
	for _, row := range b {
		for _, val := range row {
			if val >= 0 {
				sum += val
			}
		}
	}
	return score(sum * curDraw), curTurn
}

func part1(boards []board, draws []int) {
	lowestTurns, score := math.MaxInt, 0
	for _, b := range boards {
		s, t := b.scoreBoard(draws)
		if t < lowestTurns {
			lowestTurns = t
			score = int(s)
		}
	}
	fmt.Println(score)
}

func part2(boards []board, draws []int) {
	highestTurns, score := math.MinInt, 0
	for _, b := range boards {
		s, t := b.scoreBoard(draws)
		fmt.Println(s, t)
		if t > highestTurns {
			highestTurns = t
			score = int(s)
		}
	}
	fmt.Println(score)
}

func main() {
	boards, draws := readInput(INPUT_FILE)
	part1(boards, draws)
	boards, draws = readInput(INPUT_FILE)
	part2(boards, draws)
}
