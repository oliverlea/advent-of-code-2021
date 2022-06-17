package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_FILE = "input"
const GRID_SIZE = 1000

type coord struct {
	x, y int
}

type vector struct {
	start coord
	end   coord
}

type grid [][]int

func parseCoord(s string) coord {
	words := strings.Split(s, ",")
	x, _ := strconv.Atoi(words[0])
	y, _ := strconv.Atoi(words[1])
	return coord{x, y}
}

func parseVector(s string) vector {
	words := strings.Split(s, " -> ")
	startString, endString := words[0], words[1]
	return vector{parseCoord(startString), parseCoord(endString)}
}

func parseInput(fn string) []vector {
	lines := utils.ReadStrings(fn)
	vectors := make([]vector, len(lines))
	for i, line := range lines {
		vectors[i] = parseVector(line)
	}
	return vectors
}

func (c *coord) add(x int, y int) {
	c.x += x
	c.y += y
}

func (v *vector) flip() {
	tmp := v.start
	v.start = v.end
	v.end = tmp
}

func (g grid) addVector(v vector, considerDiag bool) {
	diag := false
	if v.start.x != v.end.x && v.start.y != v.end.y {
		diag = true
		if !considerDiag {
			return
		}
	}
	dy := v.end.y - v.start.y
	dx := v.end.x - v.start.x

	// Reverse vectors that are going backwards or up
	if !diag && (dy < 0 || dx < 0) {
		if dx < 0 {
			dx *= -1
		}
		if dy < 0 {
			dy *= -1
		}
		v.flip()
	}

	cur := coord{v.start.x, v.start.y}
	g[cur.y][cur.x]++
	if diag {
		// dx and dy will be the same since x = y
		xInc, yInc := 1, 1
		if v.end.x-v.start.x < 0 {
			xInc = -1
		}
		if v.end.y-v.start.y < 0 {
			yInc = -1
		}

		for dx *= xInc; dx > 0; dx-- {
			cur.add(xInc, yInc)
			g[cur.y][cur.x]++
		}
	} else {
		// Straight line
		for ; dx > 0; dx-- {
			cur.add(1, 0)
			g[cur.y][cur.x]++
		}
		for ; dy > 0; dy-- {
			cur.add(0, 1)
			g[cur.y][cur.x]++
		}
	}
}

func eval(vectors []vector, considerDiag bool) {
	grid := make(grid, GRID_SIZE)
	for i := range grid {
		grid[i] = make([]int, GRID_SIZE)
	}

	for _, v := range vectors {
		grid.addVector(v, considerDiag)
	}

	overlap := 0
	for _, row := range grid {
		for _, val := range row {
			if val > 1 {
				overlap++
			}
		}
	}
	fmt.Println(overlap)
}

func main() {
	vectors := parseInput(INPUT_FILE)
	eval(vectors, false)
	vectors = parseInput(INPUT_FILE)
	eval(vectors, true)
}
