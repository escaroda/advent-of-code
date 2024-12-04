package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	print("--- Day 1: Historian Hysteria ---\n", "https://adventofcode.com/2024/day/1\n\n")
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

func getLists() [2][]int {
	data := getData()
	lists := [2][]int{}
	for _, line := range strings.Split(strings.TrimRight(string(data), "\n"), "\n") {
		digits := strings.Fields(line)

		if len(digits) != 2 {
			log.Fatal("There should be 2 lists only")
		}

		for i, digit := range digits {
			num, err := strconv.ParseUint(digit, 10, 32)
			check(err)
			lists[i] = append(lists[i], int(num))
		}
	}
	return lists
}

func part1() {
	lists := getLists()
	for _, list := range lists {
		sort.Ints(list)
	}

	total := 0
	for i := range len(lists[0]) {
		diff := lists[0][i] - lists[1][i]
		if diff < 0 {
			total -= diff
		} else {
			total += diff
		}
	}

	println("Part 1:", total)
}

func part2() {
	lists := getLists()
	dict := make(map[int]int)
	for _, num := range lists[1] {
		dict[num] += 1
	}
	score := 0
	for _, num := range lists[0] {
		score += num * dict[num]
	}

	println("Part 2:", score)
}
