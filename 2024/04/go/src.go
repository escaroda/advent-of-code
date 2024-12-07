package main

import (
	"log"
	"os"
	"strings"
)

// Assertion: This code works with XMAS and MAS only
var t = "XMAS"

func main() {
	print("--- Day 4: Ceres Search ---\n", "https://adventofcode.com/2024/day/4\n\n")
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

func getGrid() (grid [][]int) {
	data := getData()
	for _, l := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		g := []int{}
		for _, v := range l {
			g = append(g, strings.IndexRune(t, v))
		}
		grid = append(grid, g)
	}
	return
}

func part1() {
	dirs := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	grid := getGrid()
	count := 0
	s := [][4]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				s = append(s, [4]int{i, j, 1, -1})
			}
		}
	}

	for len(s) > 0 {
		var c [4]int
		c, s = s[len(s)-1], s[:len(s)-1]
		if c[2] == len(t) {
			count += 1
			continue
		}

		for a, d := range dirs {
			if c[3] > -1 && c[3] != a {
				continue
			}
			i, j := c[0]+d[0], c[1]+d[1]
			if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[i]) && grid[i][j] == c[2] {
				s = append(s, [4]int{i, j, c[2] + 1, a})
			}
		}
	}

	println("Part 1:", count)
}

func part2() {
	count := 0
	grid := getGrid()
	s := [][2]int{}
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 2 {
				s = append(s, [2]int{i, j})
			}
		}
	}

	for _, a := range s {
		i, j := a[0], a[1]
		m1, s1 := grid[i-1][j-1], grid[i+1][j+1]
		m2, s2 := grid[i+1][j-1], grid[i-1][j+1]
		if m1 != s1 && m2 != s2 && (m1 == 1 || m1 == 3) && (s1 == 1 || s1 == 3) && (m2 == 1 || m2 == 3) && (s2 == 1 || s2 == 3) {
			count += 1
		}
	}

	println("Part 2:", count)
}
