package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

func part1(input string) int {
	stones := common.StringsToNumbers(strings.Split(input, " "))

	for i := 0; i < 25; i++ {
		newStones := []int{}
		for _, stone := range stones {
			strStone := strconv.Itoa(stone)
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strStone)%2 == 0 {
				leftStone, _ := strconv.Atoi(strStone[:len(strStone)/2])
				rightStone, _ := strconv.Atoi(strStone[len(strStone)/2:])
				newStones = append(newStones, leftStone, rightStone)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}

	return len(stones)
}

func part2(input string) int {
	return 0
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(11)
	fmt.Println("DAY 11")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
