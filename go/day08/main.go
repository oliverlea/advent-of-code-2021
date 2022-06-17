package main

import (
	"advent/utils"
	"fmt"
	"log"
	"reflect"
	"strings"
)

const INPUT_FILE = "input"

type display map[rune]bool

const TOP_MID = 'a'
const TOP_LEFT = 'b'
const TOP_RIGHT = 'c'
const MID_MID = 'd'
const BOT_LEFT = 'e'
const BOT_RIGHT = 'f'
const BOT_MID = 'g'

type signal struct {
	input  []display
	output []display
}

func (d display) add(other display) display {
	newD := make(display)
	for k, _ := range d {
		newD[k] = true
	}
	for k, _ := range other {
		newD[k] = true
	}
	return newD
}

func (d display) addRune(r rune) display {
	newD := make(display)
	for k, _ := range d {
		newD[k] = true
	}
	newD[r] = true
	return newD
}

func (d display) subtract(other display) display {
	newD := make(display)
	for k, _ := range d {
		if _, exists := other[k]; !exists {
			newD[k] = true
		}
	}
	return newD
}

func (d display) subtractRune(r rune) display {
	newD := make(display)
	for k, _ := range d {
		if k != r {
			newD[k] = true
		}
	}
	return newD
}

// func (d display) translate(m map[rune]rune) display {
// 	newD := make(display)
// 	for k, _ := range d {
// 		newD[m[k]] = true
// 	}
// 	return newD
// }

func processWords(ws []string) []display {
	ds := make([]display, len(ws))
	for i, w := range ws {
		d := make(display)
		for _, r := range w {
			d[r] = true
		}
		ds[i] = d
	}
	return ds
}

func parseSignal(line string) signal {
	parts := strings.Split(line, " | ")
	inputs, outputs := processWords(strings.Split(parts[0], " ")), processWords(strings.Split(parts[1], " "))
	return signal{inputs, outputs}
}

func (d display) categorize() int {
	switch len(d) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	default:
		return -1
	}
}

func (d display) anyValue() rune {
	for x, _ := range d {
		return x
	}
	log.Fatal("No key in set")
	return ' '
}

func part1(signals []signal) {
	total := 0
	for _, s := range signals {
		for _, out := range s.output {
			if out.categorize() >= 0 {
				total += 1
			}
		}
	}
	fmt.Println(total)
}

func supersetMatch(ds []display, target display, length int) display {
	for _, d := range ds {
		if len(d) == length && len(d.subtract(target)) > 0 && len(target.subtract(d)) == 0 {
			return d
		}
	}
	log.Fatal("Could not find match")
	return nil
}

func part2(signals []signal) {
	sum := 0
	for _, s := range signals {
		signalNumber := make(map[int]display)
		for _, input := range s.input {
			if mapping := input.categorize(); mapping >= 0 {
				signalNumber[mapping] = input
			}
		}
		wireMap := make(map[rune]rune)
		// Top mid
		wireMap[TOP_MID] = signalNumber[7].subtract(signalNumber[1]).anyValue()
		// Infer 9, bot mid and bot left
		signalNumber[9] = supersetMatch(s.input, signalNumber[4].addRune(wireMap[TOP_MID]), 6)
		wireMap[BOT_MID] = signalNumber[9].subtract(signalNumber[4]).subtractRune(wireMap[TOP_MID]).anyValue()
		wireMap[BOT_LEFT] = signalNumber[8].subtract(signalNumber[9]).anyValue()
		// Currently know: 1, 4, 7, 8, 9 + 'top mid', 'bot mid'
		// Infer 6
		signalNumber[6] = supersetMatch(s.input, signalNumber[9].subtract(signalNumber[1]).addRune(wireMap[BOT_LEFT]), 6)
		// Infer top-right
		wireMap[TOP_RIGHT] = signalNumber[8].subtract(signalNumber[6]).anyValue()
		// Infer 5
		signalNumber[5] = signalNumber[6].subtractRune(wireMap[BOT_LEFT])
		// Currently know: 1, 4, 5, 6, 7, 8, 9
		wireMap[BOT_RIGHT] = signalNumber[1].subtractRune(TOP_RIGHT).anyValue()
		// Infer 0
		signalNumber[0] = supersetMatch(s.input, signalNumber[8].subtract(signalNumber[8]).addRune(wireMap[BOT_LEFT]).addRune(wireMap[TOP_RIGHT]), 6)
		wireMap[MID_MID] = signalNumber[8].subtract(signalNumber[0]).anyValue()
		wireMap[TOP_LEFT] = signalNumber[4].subtract(signalNumber[1]).subtractRune(wireMap[MID_MID]).anyValue()
		signalNumber[2] = supersetMatch(s.input, signalNumber[8].subtract(signalNumber[5]), 5)
		signalNumber[3] = signalNumber[8].subtractRune(wireMap[BOT_LEFT]).subtractRune(wireMap[TOP_LEFT])

		values := make([]int, 0, len(s.output))
		for _, output := range s.output {
			found := false
			for value, d := range signalNumber {
				if reflect.DeepEqual(output, d) {
					values = append(values, value)
					found = true
					break
				}
			}
			if !found {
				log.Fatal("No value found")
			}
		}
		res, op := 0, 1
		for i := len(values) - 1; i >= 0; i-- {
			res += op * values[i]
			op *= 10
		}
		sum += res
	}
	fmt.Println(sum)
}

func main() {
	lines := utils.ReadStrings(INPUT_FILE)
	signals := make([]signal, len(lines))
	for i, l := range lines {
		signals[i] = parseSignal(l)
	}

	part1(signals)
	part2(signals)
}
