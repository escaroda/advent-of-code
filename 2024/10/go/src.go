package main

import (
	"log"
	"os"
	"strings"
)

var dirs = [...][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down, left

func main() {
	print("--- Day 10: Hoof It ---\n", "https://adventofcode.com/2024/day/10\n\n")
	part1()
	part2()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getData() []byte {
	data, err := os.ReadFile("../input.txt")
	check(err)
	return data
}

type point struct {
	v int
	m map[int]bool
	q int
}

func getGrid() [][]point {
	g := [][]point{}
	for _, l := range strings.Split(strings.TrimRight(string(getData()), "\n"), "\n") {
		r := []point{}
		for _, d := range l {
			n := int(d - '0')
			m := make(map[int]bool)
			r = append(r, point{n, m, 0})
		}
		g = append(g, r)
	}
	return g
}

func propagate(s map[int]bool, t map[int]bool) {
	for k, v := range s {
		t[k] = v
	}
}

func process(g [][]point) {
	s := [][2]int{} // stack
	for i, r := range g {
		for j, p := range r {
			if p.v == 9 {
				g[i][j].m[i<<16|j] = true // bitmask to store two 16bit ints in one 32bit (assume topographic map width & height fit 16bit)
				s = append(s, [2]int{i, j})
			}
		}
	}

	for len(s) > 0 {
		p := s[len(s)-1] // get last element
		s = s[:len(s)-1] // pop last element
		n := g[p[0]][p[1]].v - 1
		for _, d := range dirs {
			i, j := p[0]+d[0], p[1]+d[1]
			if i < 0 || i >= len(g) || j < 0 || j >= len(g[0]) || g[i][j].v != n {
				continue
			}
			propagate(g[p[0]][p[1]].m, g[i][j].m) // 9-height positions reachable from that point
			g[i][j].q += 1                        // distinct hiking trails
			if n > 0 {
				s = append(s, [2]int{i, j})
			}
		}
	}
}

func sum(g [][]point) (int, int) {
	s1, s2 := 0, 0
	for _, r := range g {
		for _, p := range r {
			if p.v == 0 {
				s1 += len(p.m)
				s2 += p.q
			}
		}
	}
	return s1, s2
}

func part1() {
	g := getGrid()
	process(g)

	res, _ := sum(g)

	println("Part 1:", res)
}

func part2() {
	g := getGrid()
	process(g)

	_, res := sum(g)

	println("Part 2:", res)
}
