package main

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const INPUT_FILE = "input"

type heightMap [][]int
type coord struct {
	x int
	y int
}

func (hm heightMap) valueAt(c coord) int {
	return hm[c.y][c.x]
}

func parseInput(input []string) heightMap {
	heights := make(heightMap, len(input))
	for i, s := range input {
		row := make([]int, len(s))
		for x, c := range strings.Split(s, "") {
			row[x], _ = strconv.Atoi(c)
		}
		heights[i] = row
	}
	return heights
}

func (hm heightMap) expand(c coord) []coord {
	results := make([]coord, 0, 4)
	if c.x != 0 {
		// Left
		results = append(results, coord{c.x - 1, c.y})
	}
	if c.y != 0 {
		// Up
		results = append(results, coord{c.x, c.y - 1})
	}
	if c.x != len(hm[0])-1 {
		// Right
		results = append(results, coord{c.x + 1, c.y})
	}
	if c.y != len(hm)-1 {
		// Down
		results = append(results, coord{c.x, c.y + 1})
	}
	return results
}

func (hm heightMap) lowPoints() []coord {
	lowPointCoords := make([]coord, 0)
	for y, row := range hm {
		for x, val := range row {
			curCoord, lowest := coord{x, y}, true
			for _, expandedCoord := range hm.expand(curCoord) {
				if val >= hm.valueAt(expandedCoord) {
					lowest = false
					break
				}
			}
			if lowest {
				lowPointCoords = append(lowPointCoords, curCoord)
			}
		}
	}
	return lowPointCoords
}

func part1(hm heightMap) {
	lowPointCoords, lowPointRiskSum := hm.lowPoints(), 0
	for _, c := range lowPointCoords {
		lowPointRiskSum += hm.valueAt(c) + 1
	}
	fmt.Println(lowPointRiskSum)
}

type coordPair struct {
	curCoord  coord
	prevCoord coord
}

func basinSizeExpand(hm heightMap, startCoord coord) int {
	var (
		q    = make([]coordPair, 0, 100)
		seen = make(map[coord]bool)
		pair coordPair
		size int = 0
	)
	q = append(q, coordPair{startCoord, coord{-1, -1}})
	seen[startCoord] = true

	for len(q) > 0 {
		pair = q[0]
		q = q[1:]

		if pair.prevCoord.x != -1 {
			if _, seenExpandedCoord := seen[pair.curCoord]; seenExpandedCoord {
				continue
			}
			curValue, _ := hm.valueAt(pair.curCoord), hm.valueAt(pair.prevCoord)
			if curValue == 9 {
				continue
			}
		}
		size++
		seen[pair.curCoord] = true
		for _, expandedCoord := range hm.expand(pair.curCoord) {
			q = append(q, coordPair{expandedCoord, pair.curCoord})
		}
	}
	return size
}

func part2(hm heightMap) {
	var (
		basinSizes     = make([]int, 0)
		lowPointCoords = hm.lowPoints()
	)
	for _, lowPoint := range lowPointCoords {
		basinSizes = append(basinSizes, basinSizeExpand(hm, lowPoint))
	}
	result := 1
	sort.Ints(basinSizes)
	for _, s := range basinSizes[len(basinSizes)-3:] {
		result *= s
	}
	fmt.Println(result)
}

func main() {
	input := utils.ReadStrings(INPUT_FILE)
	heights := parseInput(input)
	part1(heights)
	part2(heights)
}
