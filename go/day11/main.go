package main

import (
	"advent/utils"
	"fmt"
)

const INPUT_FILE = "input"

type grid [][]int
type coord struct {
	x int
	y int
}

func (c coord) expand() []coord {
	return []coord{
		{c.x - 1, c.y - 1}, // Top left
		{c.x, c.y - 1},     // Top mid
		{c.x + 1, c.y - 1}, // Top right
		{c.x - 1, c.y},     // Left
		{c.x + 1, c.y},     // Right
		{c.x - 1, c.y + 1}, // Bot left
		{c.x, c.y + 1},     // Bot mid
		{c.x + 1, c.y + 1}, // Bot right
	}
}

func (g grid) step() int {
	// Increment the energy and keep track of anything about to flash
	flashable := make(map[coord]bool)
	for y, row := range g {
		for x, val := range row {
			g[y][x] = val + 1
			if val+1 == 10 {
				flashable[coord{x, y}] = true
			}
		}
	}

	flashed := make(map[coord]bool)
	// For each flash, increase the adjacent coords by 1
	for len(flashable) > 0 {
		var curCoord coord
		for c := range flashable {
			// Take only a single coord per iteration of the outer loop
			curCoord = c
			break
		}
		for _, exp := range curCoord.expand() {
			if exp.y >= 0 && exp.y < len(g) && exp.x >= 0 && exp.x < len(g[exp.y]) {
				g[exp.y][exp.x]++
				cur := g[exp.y][exp.x]
				if _, hasFlashed := flashed[exp]; !hasFlashed && cur >= 10 {
					flashable[coord{exp.x, exp.y}] = true
				}
			}
		}
		flashed[curCoord] = true
		delete(flashable, curCoord)
	}

	totalFlashes := 0
	for y, row := range g {
		for x, val := range row {
			if val >= 10 {
				totalFlashes++
				g[y][x] = 0
			}
		}
	}
	return totalFlashes
}

func (g grid) simulate(steps int) int {
	total := 0
	for i := 0; i < steps; i++ {
		total += g.step()
	}
	return total
}

func (g grid) findSync() int {
	targetFlashes := len(g) * len(g[0])
	curStep := 0
	for {
		curFlashes := g.step()
		curStep++
		if curFlashes == targetFlashes {
			return curStep
		}
	}
}

func part1(g grid) {
	flashes := g.simulate(100)
	fmt.Println(flashes)
}

func part2(g grid) {
	fmt.Println(g.findSync())
}

func createGrid() grid {
	var (
		lines = utils.ReadStrings(INPUT_FILE)
		g     = make(grid, len(lines))
	)
	for y, l := range lines {
		g[y] = make([]int, len(l))
		for x, r := range l {
			g[y][x] = int(r - '0')
		}
	}
	return g
}

func main() {
	part1(createGrid())
	part2(createGrid())
}
