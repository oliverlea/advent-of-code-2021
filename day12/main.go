package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

const INPUT_FILE = "input"

type node string
type edge struct {
	from, to node
}
type graph map[node][]node

var (
	START = node("start")
	END   = node("end")
)

func (n node) isSmall() bool {
	return n[0] >= 'a'
}

func parseEdge(s string) edge {
	splitString := strings.Split(s, "-")
	return edge{node(splitString[0]), node(splitString[1])}
}

func constructGraph(edges []edge) graph {
	g := make(map[node][]node)
	for _, e := range edges {
		if _, exists := g[e.from]; !exists {
			g[e.from] = make([]node, 0)
		}
		if _, exists := g[e.to]; !exists {
			g[e.to] = make([]node, 0)
		}
		g[e.from] = append(g[e.from], e.to)
		g[e.to] = append(g[e.to], e.from)
	}
	return g
}

func bfsR(g graph, smallDoubleVisitAllowance int, cur node, target node, curPath []node) [][]node {
	// Base case. It's small and we've seen the node before
	downStream := make([][]node, 0)
	if cur == target {
		return append(downStream, append(curPath, cur))
	}
	curPath = append(curPath, cur)
	if cur.isSmall() {
		if cur == START && len(curPath) > 1 {
			return downStream
		}
		smallVisits := make(map[node]int)
		smallVisitsAbove1 := 0
		// If we've seen a small cave too many times, return
		for _, n := range curPath {
			if !n.isSmall() || cur == START || cur == END {
				continue
			}
			cur, exists := smallVisits[n]
			if exists {
				smallVisits[n]++
			} else {
				smallVisits[n] = 1
			}
			if cur+1 > 1 {
				smallVisitsAbove1++
			}
			if smallVisitsAbove1 > smallDoubleVisitAllowance {
				return downStream
			}
		}
	}

	// Recursive case
	nextNodes, hasNext := g[cur]
	if !hasNext {
		return downStream
	}

	for _, next := range nextNodes {
		r := bfsR(g, smallDoubleVisitAllowance, next, target, curPath)
		if len(r) > 0 {
			downStream = append(downStream, r...)
		}
	}
	return downStream
}

func bfs(g graph, smallVisits int, start node, target node) [][]node {
	return bfsR(g, smallVisits, start, target, make([]node, 0))
}

func part1(g graph) {
	paths := bfs(g, 0, START, END)
	fmt.Println(len(paths))
}

func part2(g graph) {
	paths := bfs(g, 1, START, END)
	fmt.Println(len(paths))
}

func main() {
	lines := utils.ReadStrings(INPUT_FILE)
	edges := make([]edge, len(lines))
	for i, l := range lines {
		edges[i] = parseEdge(l)
	}
	g := constructGraph(edges)
	part1(g)
	part2(g)
}
