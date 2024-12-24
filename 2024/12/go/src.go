package main

import (
	"log"
	"os"
	"sort"
	"strings"
)

var dirs = [...][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down, left

func main() {
	print("--- Day 12: Garden Groups ---\n", "https://adventofcode.com/2024/day/12\n\n")
	part1()
	part2()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func data() []byte {
	data, err := os.ReadFile("../input.txt")
	check(err)
	return data
}

func out(g [][][2]int, i int, j int) bool {
	return i < 0 || i >= len(g) || j < 0 || j >= len(g[0])
}

func grid() [][][2]int {
	g := [][][2]int{}
	for _, l := range strings.Split(strings.TrimRight(string(data()), "\n"), "\n") {
		r := [][2]int{}
		for _, v := range l {
			r = append(r, [2]int{int(v), 0})
		}
		g = append(g, r)
	}
	return g
}

func part1() {
	sum := 0
	g := grid()
	for i, r := range g {
		for j, v := range r {
			if v[1] == 1 {
				continue
			}
			s := [][2]int{{i, j}}
			g[i][j][1] = 1
			a, p := 1, 0
			for len(s) > 0 {
				c := s[len(s)-1]
				s = s[:len(s)-1]
				n := [][2]int{}
				pp := 0
				for _, d := range dirs {
					ii, jj := c[0]+d[0], c[1]+d[1]
					if out(g, ii, jj) {
						continue
					}
					if g[ii][jj][0] == v[0] {
						pp += 1
						if g[ii][jj][1] == 0 {
							g[ii][jj][1] = 1
							n = append(n, [2]int{ii, jj})
						}

					}
				}
				a += len(n)
				p += len(dirs) - pp
				s = append(s, n...)
			}
			sum += a * p
		}
	}

	println("Part 1:", sum)
}

func sequences(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			count += 1
		}
	}
	return count
}

func sides(ms [4]map[int]map[int]bool) int {
	res := 0
	for _, vh := range ms {
		for _, m := range vh {
			ks := []int{}
			for k := range m {
				ks = append(ks, k)
			}
			sort.Ints(ks)
			res += sequences(ks)
		}
	}
	return res
}

func insert(m map[int]map[int]bool, k, v int) {
	_, ok := m[k]
	if !ok {
		m[k] = make(map[int]bool)
	}
	m[k][v] = true
}

func part2() {
	sum := 0
	g := grid()
	for i, r := range g {
		for j, v := range r {
			if v[1] == 1 {
				continue
			}
			s := [][2]int{{i, j}}
			g[i][j][1] = 1
			a, p := 1, 0
			ms := [4]map[int]map[int]bool{ // map for each edge crossing - move up, right, down, left
				make(map[int]map[int]bool),
				make(map[int]map[int]bool),
				make(map[int]map[int]bool),
				make(map[int]map[int]bool),
			}
			for len(s) > 0 {
				c := s[len(s)-1]
				s = s[:len(s)-1]
				n := [][2]int{}
				for o, d := range dirs {
					i, j := c[0]+d[0], c[1]+d[1]
					if out(g, i, j) || g[i][j][0] != v[0] {
						if d[0] == 0 {
							insert(ms[o], j, i) // crossing vertical edge
						} else {
							insert(ms[o], i, j) // crossing horizontal edge
						}
					} else if g[i][j][1] == 0 {
						g[i][j][1] = 1
						n = append(n, [2]int{i, j})
					}
				}
				a += len(n)
				s = append(s, n...)
			}
			if a > 2 {
				p = sides(ms)
			} else {
				p = 4
			}
			sum += a * p
		}
	}

	println("Part 2:", sum)
}
