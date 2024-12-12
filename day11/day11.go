package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

type stoneIter struct {
	stone int
	iter  int
}

var memo = make(map[stoneIter]int)

func blink(stone int, iter int) int {
	if iter == 0 {
		return 1
	}

	memoed := memo[stoneIter{stone, iter}]
	if memoed != 0 {
		return memoed
	}

	total := 0

	strStone := strconv.Itoa(stone)
	if stone == 0 {
		total += blink(1, iter-1)
	} else if len(strStone)%2 == 0 {
		leftStone, _ := strconv.Atoi(strStone[:len(strStone)/2])
		rightStone, _ := strconv.Atoi(strStone[len(strStone)/2:])
		total += blink(leftStone, iter-1) + blink(rightStone, iter-1)
	} else {
		total = blink(stone*2024, iter-1)
	}

	memo[stoneIter{stone, iter}] = total

	return total
}

func process(stones []int, iter int) int {
	var total = 0
	for _, num := range stones {
		total += blink(num, iter)
	}
	return total
}

func part1(input string) int {
	stones := common.StringsToNumbers(strings.Split(input, " "))
	return process(stones, 25)
}

func part2(input string) int {
	stones := common.StringsToNumbers(strings.Split(input, " "))
	return process(stones, 75)
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(11)
	fmt.Println("DAY 11")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
