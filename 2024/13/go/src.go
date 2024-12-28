package main

import (
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var digits = regexp.MustCompile("\\d+")

const shift = 1e13

type Point struct {
	x, y int
}

type Machine struct {
	a, b, p Point
}

func (m Machine) Cost() int {
	ax, ay, bx, by, px, py := float64(m.a.x), float64(m.a.y), float64(m.b.x), float64(m.b.y), float64(m.p.x), float64(m.p.y)
	i := (by*px - bx*py) / (ax*by - ay*bx)
	j := (px - ax*i) / bx
	if i != math.Trunc(i) || j != math.Trunc(j) {
		return 0
	}
	return int(i)*3 + int(j)
}

func (m *Machine) ShiftPrize() {
	(*m).p.x += shift
	(*m).p.y += shift
}

func main() {
	print("--- Day 13: Claw Contraption ---\n", "https://adventofcode.com/2024/day/13\n\n")
	part1()
	part2()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func data() []byte {
	data, err := os.ReadFile("../input.txt")
	check(err)
	return data
}

func machines() []Machine {
	res := []Machine{}
	for _, s := range strings.Split(strings.TrimRight(string(data()), "\n"), "\n\n") {
		v := [6]int{}
		for i, d := range digits.FindAllString(s, -1) {
			n, err := strconv.Atoi(d)
			check(err)
			v[i] = n
		}
		res = append(res, Machine{Point{v[0], v[1]}, Point{v[2], v[3]}, Point{v[4], v[5]}})
	}
	return res
}

func part1() {
	sum := 0
NEXT_MACHINE:
	for _, m := range machines() {
		for i := range 101 { // brute force
			for j := range 101 {
				if m.a.x*i+m.b.x*j != m.p.x || m.a.y*i+m.b.y*j != m.p.y {
					continue
				}
				sum += i*3 + j
				continue NEXT_MACHINE
			}
		}
	}

	println("Part 1:", sum)
}

func part2() {
	sum := 0
	for _, m := range machines() {
		m.ShiftPrize()
		sum += m.Cost()
	}

	println("Part 2:", sum)
}
