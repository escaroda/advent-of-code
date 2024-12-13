package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	print("--- Day 6: Guard Gallivant ---\n", "https://adventofcode.com/2024/day/6\n\n")
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

func getGridAndPos() ([][]int, [2]int) {
	grid := [][]int{}
	pos := [2]int{}
	data := getData()
	for i, l := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		g := []int{}
		for j, v := range l {
			var c int
			switch v {
			case '#':
				c = 9
			case '^':
				pos = [2]int{i, j}
			}
			g = append(g, c)
		}
		grid = append(grid, g)
	}
	return grid, pos
}

func copyGrid(g [][]int) (c [][]int) {
	for i := range g {
		c = append(c, []int{})
		for j := range g[i] {
			v := 0
			if g[i][j] == 9 {
				v = 9
			}
			c[i] = append(c[i], v)
		}
	}
	return
}

func turnRight(p *[2]int) {
	temp := p[0]
	p[0] = p[1]
	p[1] = -temp
}

func walk(grid [][]int, pos [2]int) {
	dir := [2]int{-1, 0}
	for {
		i, j := pos[0]+dir[0], pos[1]+dir[1]
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[i])-1 {
			break
		}
		switch grid[i][j] {
		case 0:
			(grid)[i][j] += 1
			pos = [2]int{i, j}
		case 1:
			pos = [2]int{i, j}
		case 9:
			turnRight(&dir)
		}
	}
}

func part1() {
	grid, pos := getGridAndPos()
	walk(grid, pos)

	count := 0
	for _, l := range grid {
		for _, v := range l {
			if v == 1 {
				count += 1
			}
		}
	}

	println("Part 1:", count)
}

func part2() {
	grid, pos := getGridAndPos()
	walk(grid, pos)
	init := pos
	count := 0
	points := [][2]int{}
	for i, l := range grid {
		for j, v := range l {
			if v != 1 || (i == pos[0] && j == pos[1]) {
				continue
			}
			points = append(points, [2]int{i, j})
		}
	}

NEXT:
	for _, p := range points {
		pos := init
		dir := [2]int{-1, 0}
		g := copyGrid(grid)
		g[p[0]][p[1]] = 9 // place obstacle
		for {
			i, j := pos[0]+dir[0], pos[1]+dir[1]
			if i < 0 || i > len(g)-1 || j < 0 || j > len(g[i])-1 {
				break
			}
			switch g[i][j] {
			case 4:
				count += 1
				continue NEXT
			case 9:
				turnRight(&dir)
			default:
				pos = [2]int{i, j}
				g[i][j] += 1
			}
		}
	}

	println("Part 2:", count)
}
