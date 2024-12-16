package main

import (
	"fmt"
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

func parseLine(line string) coord {
	matches := lineRegex.FindAllStringSubmatch(line, -1)
	return coord{common.StrToInt(matches[0][1]), common.StrToInt(matches[0][2])}
}

func parseInput(input string, extraPrize int) []machine {
	machineDefs := strings.Split(input, "\n\n")
	machines := make([]machine, len(machineDefs))
	for i, def := range machineDefs {
		lines := strings.Split(def, "\n")
		prize := parseLine(lines[2])
		machines[i] = machine{
			a: parseLine(lines[0]),
			b: parseLine(lines[1]),
			prize: coord{
				x: prize.x + extraPrize,
				y: prize.y + extraPrize,
			},
		}
	}
	return machines
}

func machineCost(m machine) int {
	b0 := (m.prize.y*m.a.x - m.a.y*m.prize.x) / (m.b.y*m.a.x - m.a.y*m.b.x)
	a0 := (m.prize.x - b0*m.b.x) / m.a.x

	if (m.a.x*a0+m.b.x*b0) == m.prize.x && (m.a.y*a0+m.b.y*b0) == m.prize.y {
		return A_COST*a0 + B_COST*b0
	} else {
		return 0
	}
}

func part1(input string) int {
	machines := parseInput(input, 0)
	totalCost := 0

	for _, machine := range machines {
		totalCost += machineCost(machine)
	}

	return totalCost
}

func part2(input string) int {
	machines := parseInput(input, 10000000000000)
	totalCost := 0

	for _, machine := range machines {
		totalCost += machineCost(machine)
	}

	return totalCost
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(13)
	fmt.Println("DAY 13")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
