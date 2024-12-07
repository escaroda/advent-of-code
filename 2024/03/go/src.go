package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

var dig = regexp.MustCompile(`\d+`)
var mul = regexp.MustCompile(`mul\(\d+,\d+\)`)
var don = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

func main() {
	print("--- Day 3: Mull It Over ---\n", "https://adventofcode.com/2024/day/3\n\n")
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

func multiply(s string) int {
	res := 1
	for _, v := range dig.FindAllString(s, -1) {
		n, err := strconv.Atoi(v)
		check(err)
		res *= n
	}
	return res
}

func part1() {
	mem := string(getData())
	sum := 0
	for _, match := range mul.FindAllString(mem, -1) {
		sum += multiply(match)
	}

	println("Part 1:", sum)
}

func part2() {
	mem := string(getData())
	e := true
	sum := 0
	for _, match := range don.FindAllString(mem, -1) {
		if match == "don't()" {
			e = false
		} else if match == "do()" {
			e = true
		} else if e {
			sum += multiply(match)
		}
	}

	println("Part 2:", sum)
}
