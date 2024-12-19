package main

import (
	"log"
	"os"
	"strings"
)

const antinode rune = '#'
const empty rune = '.'

func main() {
	print("--- Day 8: Resonant Collinearity ---\n", "https://adventofcode.com/2024/day/8\n\n")
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

func getMap() [][][2]rune {
	m := [][][2]rune{}
	for _, line := range strings.Split(strings.TrimRight(string(getData()), "\n"), "\n") {
		l := [][2]rune{}
		for _, c := range line {
			l = append(l, [2]rune{c, empty})
		}
		m = append(m, l)
	}
	return m
}

func find(m [][][2]rune, i int, j int) [][2]int {
	res := [][2]int{}
	c := m[i][j]
	for ; i < len(m); i++ {
		for ; j < len(m[0]); j++ {
			if m[i][j] == c {
				res = append(res, [2]int{i, j})
			}
		}
	}
	return res
}

func mark(m [][][2]rune, ps [][2]int) {
	for _, p := range ps {
		i, j := p[0], p[1]
		if i < 0 || i >= len(m) || j < 0 || j >= len(m[0]) {
			continue
		}
		m[i][j][1] = antinode
	}
}

func part1() {
	m := getMap()
	n := make(map[rune][][2]int)
	for i, l := range m {
		for j, c := range l {
			if c[0] == empty {
				continue
			}
			n[c[0]] = append(n[c[0]], [2]int{i, j})
		}
	}

	for _, ps := range n {
		for i := 0; i < len(ps)-1; i++ {
			for j := i + 1; j < len(ps); j++ {
				p1, p2 := ps[i], ps[j]
				di, dj := p2[0]-p1[0], p2[1]-p1[1]
				mark(m, [][2]int{{p1[0] - di, p1[1] - dj}, {p2[0] + di, p2[1] + dj}})
			}
		}
	}

	count := 0
	for _, l := range m {
		for _, c := range l {
			if c[1] == antinode {
				count += 1
			}
		}
	}

	println("Part 1:", count)
}

func part2() {
	m := getMap()
	h, w := len(m), len(m[0])
	n := make(map[rune][][2]int)
	for i, l := range m {
		for j, c := range l {
			if c[0] == empty {
				continue
			}
			n[c[0]] = append(n[c[0]], [2]int{i, j})
		}
	}

	for _, ps := range n {
		for i := 0; i < len(ps)-1; i++ {
			for j := i + 1; j < len(ps); j++ {
				p1, p2 := ps[i], ps[j]
				di, dj := p2[0]-p1[0], p2[1]-p1[1]
				for ii, jj := p1[0], p1[1]; ii >= 0 && jj >= 0 && ii < h && jj < w; ii, jj = ii-di, jj-dj {
					m[ii][jj][1] = antinode
				}
				for ii, jj := p2[0], p2[1]; ii >= 0 && jj >= 0 && ii < h && jj < w; ii, jj = ii+di, jj+dj {
					m[ii][jj][1] = antinode
				}
			}
		}
	}

	count := 0
	for _, l := range m {
		for _, c := range l {
			if c[1] == antinode {
				count += 1
			}
		}
	}

	println("Part 2:", count)
}
