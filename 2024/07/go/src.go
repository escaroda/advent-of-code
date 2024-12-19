package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	print("--- Day 7: Bridge Repair ---\n", "https://adventofcode.com/2024/day/7\n\n")
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

func getLines() [][]int {
	lines := [][]int{}
	for _, line := range strings.Split(strings.TrimRight(string(getData()), "\n"), "\n") {
		l := []int{}
		for _, digit := range strings.Fields(line) {
			num, err := strconv.Atoi(strings.Trim(digit, ":"))
			check(err)
			l = append(l, num)
		}
		lines = append(lines, l)
	}
	return lines
}

func concat(a, b int) int {
	d := strconv.Itoa(a) + strconv.Itoa(b)
	n, err := strconv.Atoi(d)
	check(err)
	return n
}

func part1() {
	sum := 0
	for _, l := range getLines() {
		vs := []int{l[1]}
		for _, n := range l[2:] {
			w := []int{}
			for _, v := range vs {
				w = append(w, v+n, v*n)
			}
			vs = w
		}

		for _, v := range vs {
			if v == l[0] {
				sum += l[0]
				break
			}
		}
	}

	println("Part 1:", sum)
}

func part2() {
	sum := 0
	for _, l := range getLines() {
		vs := []int{l[1]}
		for _, n := range l[2:] {
			w := []int{}
			for _, v := range vs {
				w = append(w, v+n, v*n, concat(v, n))
			}
			vs = w
		}

		for _, v := range vs {
			if v == l[0] {
				sum += l[0]
				break
			}
		}
	}

	println("Part 2:", sum)
}
