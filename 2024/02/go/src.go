package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	print("--- Day 2: Red-Nosed Reports ---\n", "https://adventofcode.com/2024/day/2\n\n")
	part1()
	part2()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getData() []byte {
	data, err := os.ReadFile("../input.txt")
	check(err)
	return data
}

func getRows() [][]int {
	data := getData()
	rows := [][]int{}
	for _, line := range strings.Split(strings.TrimRight(string(data), "\n"), "\n") {
		row := []int{}
		for _, digit := range strings.Fields(line) {
			num, err := strconv.ParseUint(digit, 10, 8)
			check(err)
			row = append(row, int(num))
		}
		rows = append(rows, row)
	}
	return rows
}

func isSafe(row []int) bool {
	inc := true
	dec := true
	a := row[0]
	for _, b := range row[1:] {
		if a == b || abs(a-b) > 3 {
			return false
		}
		if a < b {
			dec = false
		} else {
			inc = false
		}
		a = b
	}
	return inc || dec
}

func tolerate(row []int, i int) (r []int) {
	r = append(r, row[:i]...)
	r = append(r, row[i+1:]...)
	return
}

func part1() {
	rows := getRows()
	sum := 0
	for _, row := range rows {
		if isSafe(row) {
			sum += 1
		}
	}

	println("Part 1:", sum)
}

func part2() {
	rows := getRows()
	sum := 0
	for _, row := range rows {
		if isSafe(row) {
			sum += 1
			continue
		}
		for i := 0; i < len(row); i++ {
			if isSafe(tolerate(row, i)) {
				sum += 1
				break
			}
		}
	}

	println("Part 2:", sum)
}
