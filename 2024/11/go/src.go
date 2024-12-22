package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	print("--- Day 11: Plutonian Pebbles ---\n", "https://adventofcode.com/2024/day/11\n\n")
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

func getStones() []uint {
	res := []uint{}
	for _, d := range strings.Fields(string(getData())) {
		n, err := strconv.Atoi(d)
		check(err)
		res = append(res, uint(n)) // assume all digits are non-negative
	}
	return res
}

// Integer power: compute a**b using binary powering algorithm
// See Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
func pow(a, b uint) uint {
	var p uint = 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func length(x uint) int {
	n := 0
	for x > 0 {
		x /= 10
		n++
	}
	return n
}

func isEven(x uint) bool {
	l := length(x)
	return l%2 == 0
}

func divided(x uint) []uint {
	l := uint(length(x) / 2)
	d := pow(10, l)
	return []uint{x / d, x % d}
}

func part1() {
	const blinks int = 25 // Let's simulate this part, knowing it won't work for larger amount of blinks
	s := getStones()
	for range blinks {
		ss := []uint{}
		for _, n := range s {
			switch {
			case n == 0:
				ss = append(ss, 1)
			case isEven(n):
				ss = append(ss, divided(n)...)
			default:
				ss = append(ss, n*2024)
			}
		}
		s = ss
	}

	println("Part 1:", len(s))
}

func part2() {
	const blinks uint = 75
	s := make(map[uint]uint)
	for _, n := range getStones() {
		s[n] = 1
	}
	for range blinks {
		ss := make(map[uint]uint)
		for k := range s {
			switch {
			case k == 0:
				ss[1] += s[0]
			case isEven(k):
				d := divided(k)
				ss[d[0]] += s[k]
				ss[d[1]] += s[k]
			default:
				ss[k*2024] += s[k]
			}

		}
		s = ss
	}

	var count uint
	for _, n := range s {
		count += n
	}

	println("Part 2:", count)
}
