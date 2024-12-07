package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	print("--- Day 5: Print Queue ---\n", "https://adventofcode.com/2024/day/5\n\n")
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

func getRulesAndUpdates(data []byte) ([][2]int, [][]int) {
	p := strings.Split(strings.Trim(string(data), "\n"), "\n\n")
	if len(p) != 2 {
		log.Fatal("Wrong input data")
	}
	rules := [][2]int{}
	for _, l := range strings.Split(p[0], "\n") {
		s := strings.Split(l, "|")
		if len(s) != 2 {
			log.Fatal("Wrong input data (rules part)")
		}
		r := [2]int{}
		for i := range s {
			v, err := strconv.Atoi(s[i])
			check(err)
			r[i] = v
		}
		rules = append(rules, r)
	}
	updates := [][]int{}
	for _, l := range strings.Split(p[1], "\n") {
		s := strings.Split(l, ",")
		if len(s)%2 == 0 {
			log.Fatal("Wrong input data (updates part)")
		}
		u := []int{}
		for i := range s {
			v, err := strconv.Atoi(s[i])
			check(err)
			u = append(u, v)
		}
		updates = append(updates, u)
	}

	return rules, updates
}

func part1() {
	sum := 0
	rules, updates := getRulesAndUpdates(getData())
NEXT:
	for _, u := range updates {
		for i := 0; i < len(u)-1; i++ {
			for j := i + 1; j < len(u); j++ {
				a, b := u[i], u[j]
				for _, r := range rules {
					if a == r[1] && b == r[0] {
						continue NEXT
					}
				}
			}
		}
		sum += u[len(u)/2]
	}

	println("Part 1:", sum)
}

func part2() {
	sum := 0
	rules, updates := getRulesAndUpdates(getData())
	inc := [][]int{}
NEXT:
	for _, u := range updates {
		for i := 0; i < len(u)-1; i++ {
			for j := i + 1; j < len(u); j++ {
				a, b := u[i], u[j]
				for _, r := range rules {
					if a == r[1] && b == r[0] {
						inc = append(inc, u)
						continue NEXT
					}
				}
			}
		}
	}

	for _, u := range inc {
		slices.SortFunc(u, func(a, b int) int {
			for _, r := range rules {
				if a == r[1] && b == r[0] {
					return 1
				} else if a == r[0] && b == r[1] {
					return -1
				}
			}
			return 0
		})
		sum += u[len(u)/2]
	}

	println("Part 2:", sum)
}
