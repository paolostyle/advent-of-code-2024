package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

func toCode(input string) []int {
	chars := strings.Split(input, "")
	return common.StringsToNumbers(chars)
}

func part1(input string) int {
	code := toCode(input)

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
	code := toCode(input)

	blocksLen := 0
	for _, val := range code {
		blocksLen += val
	}

	blocks := make([]int, blocksLen)
	recs := make(map[int]int)
	i := 0

	for j := 0; j < len(code); j++ {
		recs[j] = i
		if j%2 == 0 {
			for k := code[j]; k > 0; k-- {
				blocks[i] = j / 2
				i++
			}
		} else {
			i += code[j]
		}
	}

	id := 1
	for j := len(code) - 1; j >= 0; j -= 2 {
		for k := id; k <= j; k += 2 {
			if code[k] >= code[j] {
				i := recs[k]
				for l := 0; l < code[j]; l++ {
					blocks[i] = j / 2
					i++
				}

				i = recs[j]
				for l := 0; l < code[j]; l++ {
					blocks[i] = 0
					i++
				}

				code[k] -= code[j]
				recs[k] += code[j]
				code[j] = 0

				if code[id] == 0 {
					id += 2
				}
				break
			}
		}
	}

	checksum := 0
	for i, val := range blocks {
		checksum += i * val
	}

	return checksum
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(9)
	fmt.Println("DAY 09")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
