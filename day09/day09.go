package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

func part1(input string) int {
	chars := strings.Split(input, "")
	code := make([]int, len(chars))
	for i, val := range chars {
		code[i], _ = strconv.Atoi(val)
	}

	blocksLen := 0
	for i, val := range code {
		if i%2 == 0 {
			blocksLen += val
		}
	}

	blocks := make([]int, blocksLen)
	i := 0
	id := len(code) - 1

blocksLoop:
	for j := 0; j < len(code); j++ {
		if j%2 == 0 {
			for k := code[j]; k > 0; k-- {
				blocks[i] = j / 2
				i++
			}
		} else {
			emptyFields := code[j]
			for k := 0; k < emptyFields; k++ {
				if code[id] == 0 {
					id -= 2
				}
				if i >= blocksLen {
					break blocksLoop
				}
				blocks[i] = id / 2
				code[id]--
				i++
			}
		}
	}

	checksum := 0
	for i, val := range blocks {
		checksum += i * val
	}

	return checksum
}

func part2(input string) int {
	return 0
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(9)
	fmt.Println("DAY 09")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
