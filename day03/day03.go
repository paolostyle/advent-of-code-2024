package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

func part1(input string) int {
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	sum := 0
	matches := pattern.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		a := common.StrToInt(match[1])
		b := common.StrToInt(match[2])
		sum += a * b
	}

	return sum
}

func part2(input string) int {
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	sum := 0
	matches := pattern.FindAllStringSubmatch(input, -1)
	ignore := false
	for _, match := range matches {
		if match[0] == "do()" {
			ignore = false
			continue
		} else if match[0] == "don't()" {
			ignore = true
		}

		if ignore {
			continue
		}

		a := common.StrToInt(match[1])
		b := common.StrToInt(match[2])
		sum += a * b
	}

	return sum
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(3)
	fmt.Println("DAY 03")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
