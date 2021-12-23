package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_FILE = "input"

type axis int

const (
	FOLD_X axis = iota
	FOLD_Y
)

type coord struct {
	x, y int
}
type fold struct {
	v int
	a axis
}

type paper struct {
	dots          map[coord]bool
	folds         []fold
	height, width int
}

func parseInput(lines []string) paper {
	var (
		dots     = make(map[coord]bool)
		folds    = make([]fold, 0)
		readDots = true
		height   int
		width    int
	)

	for _, l := range lines {
		if len(strings.TrimSpace(l)) == 0 {
			readDots = false
			continue
		}
		if readDots {
			splitL := strings.Split(l, ",")
			var x, y int
			x, _ = strconv.Atoi(splitL[0])
			if x > width {
				width = x
			}
			y, _ = strconv.Atoi(splitL[1])
			if y > height {
				height = y
			}
			dots[coord{x, y}] = true
		} else {
			splitL := strings.Split(l, " ")
			splitL = strings.Split(splitL[2], "=")
			var a axis = FOLD_X
			if splitL[0] == "y" {
				a = FOLD_Y
			}
			v, _ := strconv.Atoi(splitL[1])
			folds = append(folds, fold{v, a})
		}
	}
	return paper{dots, folds, height + 1, width + 1}
}

func (f fold) apply(p *paper) {
	var (
		newDots = make(map[coord]bool)
	)
	for c := range p.dots {
		var dx, dy int
		if f.a == FOLD_X {
			if c.x > f.v {
				dx = c.x - f.v
				dy = 0
				newDots[coord{f.v - dx, c.y}] = true
			} else {
				newDots[c] = true
			}
		} else {
			if c.y > f.v {
				dx = 0
				dy = c.y - f.v
				newDots[coord{c.x, f.v - dy}] = true
			} else {
				newDots[c] = true
			}
		}
	}
	p.dots = newDots
	if f.a == FOLD_X {
		p.width = p.width / 2
	} else {
		p.height = p.height / 2
	}
}

func part1(p paper) {
	f := p.folds[0]
	f.apply(&p)
	fmt.Println(len(p.dots))
}

func part2(p paper) {
	for _, f := range p.folds {
		f.apply(&p)
	}
	for y := 0; y < p.height; y++ {
		line := make([]rune, p.width)
		for x := 0; x < p.width; x++ {
			c := coord{x, y}
			_, exists := p.dots[c]
			if exists {
				line[x] = '#'
			} else {
				line[x] = '.'
			}
		}
		fmt.Println(string(line))
	}
}

func main() {
	lines := utils.ReadStrings(INPUT_FILE)
	p := parseInput(lines)
	part1(p)
	lines = utils.ReadStrings(INPUT_FILE)
	p = parseInput(lines)
	part2(p)

}
