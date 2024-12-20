package main

import (
	"log"
	"os"
	"strings"
)

const empty int = -1

func main() {
	print("--- Day 9: Disk Fragmenter ---\n", "https://adventofcode.com/2024/day/9\n\n")
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

func expanded() []int {
	disk := []int{}
	for i, c := range strings.TrimRight(string(getData()), "\n") {
		n := int(c - '0')
		id := i / 2
		if i%2 == 1 { // if empty
			id = empty
		}
		for range n {
			disk = append(disk, id)
		}
	}
	return disk
}

func sum(disk []int) int {
	sum := 0
	for i, n := range disk {
		if n == empty {
			continue
		}
		sum += i * n
		i++
	}
	return sum
}

func part1() {
	disk := expanded()

	for i, j := 0, len(disk)-1; i < j; i++ {
		if disk[i] == empty {
			for j > -1 && i < j && disk[j] == empty {
				j--
			}
			disk[i], disk[j] = disk[j], disk[i]
		}
	}

	res := sum(disk)

	println("Part 1:", res)
}

func part2() {
	disk := expanded()

	j := len(disk) - 1
	for j > -1 {
		for j > -1 && disk[j] == empty {
			j--
		}
		r := j
		id := disk[j]
		for j > -1 && disk[j] == id {
			j--
		}
		l := j
		for i := 0; i <= j; i++ {
			if disk[i] != empty {
				continue
			}
			m := i
			for disk[i] == empty && i <= j {
				i++
			}
			n := i
			if n-m >= r-l { // if space available
				for d := range r - l {
					disk[m+d], disk[l+d+1] = disk[l+d+1], disk[m+d]
				}
				break
			}
		}
	}

	res := sum(disk)

	println("Part 2:", res)
}
