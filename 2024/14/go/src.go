package main

import (
	"log"
	"os"
)

func main() {
	print("--- Day 14: Restroom Redoubt ---\n", "https://adventofcode.com/2024/day/14\n\n")
	// part1()
	// part2()
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
