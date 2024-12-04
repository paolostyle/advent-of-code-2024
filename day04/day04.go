package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/paolostyle/advent-of-code-2024/common"
)

func part1(input string) int {
	rows := strings.Split(input, "\n")

	if len(rows) != len(rows[0]) {
		log.Fatal("Invalid input")
	}

	size := len(rows)

	cols := make([]string, 0)
	for i := 0; i < size; i++ {
		var sb strings.Builder
		for j := 0; j < size; j++ {
			sb.WriteByte(rows[j][i])
		}
		cols = append(cols, sb.String())
	}

	diag := make([]string, 0)
	for i := 0; i <= size-4; i++ {
		var rdown strings.Builder
		var ldown strings.Builder
		var rup strings.Builder
		var lup strings.Builder

		for j := 0; j < size; j++ {
			if i+j >= size {
				break
			}
			rup.WriteByte(rows[j][i+j])
			lup.WriteByte(rows[j][size-i-j-1])
			if i != 0 {
				rdown.WriteByte(rows[i+j][j])
				ldown.WriteByte(rows[i+j][size-j-1])
			}
		}
		diag = append(diag, rup.String(), lup.String())
		if i != 0 {
			diag = append(diag, rdown.String(), ldown.String())
		}
	}

	all := make([]string, 0)
	all = append(all, rows...)
	all = append(all, cols...)
	all = append(all, diag...)

	total := 0
	for _, s := range all {
		total += strings.Count(s, "XMAS")
		total += strings.Count(s, "SAMX")
	}

	return total
}

func part2(input string) int {
	rows := strings.Split(input, "\n")
	size := len(rows)

	total := 0

	for i := 0; i <= size-3; i++ {
		for j := 0; j <= size-3; j++ {
			var toRightSb strings.Builder
			var toLeftSb strings.Builder

			toRightSb.WriteByte(rows[i][j])
			toRightSb.WriteByte(rows[i+1][j+1])
			toRightSb.WriteByte(rows[i+2][j+2])
			toRight := toRightSb.String()

			toLeftSb.WriteByte(rows[i][j+2])
			toLeftSb.WriteByte(rows[i+1][j+1])
			toLeftSb.WriteByte(rows[i+2][j])
			toLeft := toLeftSb.String()

			if (toRight == "MAS" || toRight == "SAM") && (toLeft == "MAS" || toLeft == "SAM") {
				total += 1
			}
		}
	}

	return total
}

func main() {
	input := common.ReadInput(4)
	fmt.Println("DAY 04")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
