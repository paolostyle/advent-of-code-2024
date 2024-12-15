package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

const A_COST = 3
const B_COST = 1

var lineRegex = regexp.MustCompile(`^.*: X[+=](\d+), Y[+=](\d+)`)

type coord struct {
	x, y int
}

type machine struct {
	a, b, prize coord
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func parseLine(line string) coord {
	matches := lineRegex.FindAllStringSubmatch(line, -1)
	return coord{common.StrToInt(matches[0][1]), common.StrToInt(matches[0][2])}
}

func parseInput(input string) []machine {
	machineDefs := strings.Split(input, "\n\n")
	machines := make([]machine, len(machineDefs))
	for i, def := range machineDefs {
		lines := strings.Split(def, "\n")
		machines[i] = machine{
			a:     parseLine(lines[0]),
			b:     parseLine(lines[1]),
			prize: parseLine(lines[2]),
		}
	}
	return machines
}

func part1(input string) int {
	machines := parseInput(input)
	totalCost := 0

	for _, machine := range machines {
		ySolutions := make(map[coord]bool)
		gcdY := gcd(machine.a.y, machine.b.y)

		initialY := -1
		for a := 0; a <= 100; a++ {
			b := machine.prize.y/machine.b.y - machine.a.y*a/machine.b.y

			if machine.a.y*a+machine.b.y*b == machine.prize.y {
				initialY = a
				break
			}
		}

		if initialY == -1 {
			continue
		}

		for a := initialY; a <= 100; a += machine.b.y / gcdY {
			b := machine.prize.y/machine.b.y - machine.a.y*a/machine.b.y
			ySolutions[coord{a, b}] = true
		}

		xSolutions := make(map[coord]bool)
		gcdX := gcd(machine.a.x, machine.b.x)

		initialX := -1
		for a := 0; a <= 100; a++ {
			b := machine.prize.x/machine.b.x - machine.a.x*a/machine.b.x

			if machine.a.x*a+machine.b.x*b == machine.prize.x {
				initialX = a
				break
			}
		}

		if initialX == -1 {
			continue
		}

		for a := initialX; a <= 100; a += machine.b.x / gcdX {
			b := machine.prize.x/machine.b.x - machine.a.x*a/machine.b.x
			xSolutions[coord{a, b}] = true
		}

		validSolutions := make(map[coord]bool)

		for ySolution := range ySolutions {
			if xSolutions[ySolution] {
				validSolutions[ySolution] = true
			}
		}

		if len(validSolutions) == 0 {
			continue
		}

		minCost := math.MaxInt32
		for soution := range validSolutions {
			cost := A_COST*soution.x + B_COST*soution.y
			if cost < minCost {
				minCost = cost
			}
		}

		totalCost += minCost
	}

	return totalCost
}

func part2(input string) int {
	return 0
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(13)
	fmt.Println("DAY 13")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

/// -3 * 86 + 7 * 37 = 1
///
