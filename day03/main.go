package main

import (
	"advent/utils"
	"fmt"
	"log"
	"math"
)

const INPUT_FILE = "input"

func binaryToInt(bs []bool) int {
	x := 0
	for i, b := range bs {
		if b {
			x += int(math.Exp2(float64(len(bs) - 1 - i)))
		}
	}
	return x
}

func part1(m [][]bool) {
	rowLen := len(m[0])
	gammaBinary := make([]bool, rowLen)
	epsBinary := make([]bool, rowLen)
	for x := 0; x < rowLen; x++ {
		zeros, ones := 0, 0
		for y := 0; y < len(m); y++ {
			switch {
			case m[y][x]:
				ones++
			default:
				zeros++
			}
		}
		curVal := false
		if ones > zeros {
			curVal = true
		}
		gammaBinary[x] = curVal
		epsBinary[x] = !curVal
	}

	fmt.Println(gammaBinary, epsBinary)
	gamma, epsilon := binaryToInt(gammaBinary), binaryToInt(epsBinary)
	fmt.Println(gamma, epsilon, gamma*epsilon)
}

func matchingCommon(m [][]bool, pickFn func([]int, []int) []int) int {
	rowLen := len(m[0])
	matchingIndexes := make([]int, len(m))
	for i := 0; i < len(m); i++ {
		matchingIndexes[i] = i
	}

	for x := 0; x < rowLen; x++ {
		if len(matchingIndexes) == 1 {
			break
		}

		trues, falses := make([]int, 0, len(m)), make([]int, 0, len(m))
		for _, index := range matchingIndexes {
			if m[index][x] {
				trues = append(trues, index)
			} else {
				falses = append(falses, index)
			}
			matchingIndexes = pickFn(trues, falses)
		}
	}

	if len(matchingIndexes) != 1 {
		log.Fatal(matchingIndexes)
	}
	return binaryToInt(m[matchingIndexes[0]])
}

func part2(m [][]bool) {
	oxygen := matchingCommon(m, func(ts []int, fs []int) []int {
		switch {
		case len(ts) < len(fs):
			return fs
		case len(ts) > len(fs):
			return ts
		default:
			return ts
		}
	})
	co2 := matchingCommon(m, func(ts []int, fs []int) []int {
		switch {
		case len(ts) < len(fs):
			return ts
		case len(ts) > len(fs):
			return fs
		default:
			return fs
		}
	})
	fmt.Println(oxygen, co2, oxygen*co2)

}

func main() {
	lines := utils.ReadStrings(INPUT_FILE)
	m := make([][]bool, len(lines))
	for i, l := range lines {
		bs := []byte(l)
		bools := make([]bool, len(bs))
		for i, b := range bs {
			v := false
			if b == '1' {
				v = true
			}
			bools[i] = v
		}
		m[i] = bools
	}

	part1(m)
	part2(m)
}
