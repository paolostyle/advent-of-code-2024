package main

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/paolostyle/advent-of-code-2024/common"
)

type pos struct {
	col int
	row int
}

func newDirection(dir rune) rune {
	switch dir {
	case 'T':
		return 'R'
	case 'R':
		return 'B'
	case 'B':
		return 'L'
	case 'L':
		return 'T'
	}
	return 'T'
}

func part1(input string) int {
	rows := strings.Split(input, "\n")

	position := pos{col: 0, row: 0}
	dir := 'T'
	visited := mapset.NewSet[pos]()

	for i, row := range rows {
		col := strings.Index(row, "^")
		if col != -1 {
			position.col = col
			position.row = i
			visited.Add(position)
			break
		}
	}

	for {
		if dir == 'T' {
			for i := position.row; i >= 0; i-- {
				if rows[i][position.col] == '#' {
					dir = newDirection(dir)
					break
				} else {
					position.row = i
					visited.Add(position)
				}
			}
			if position.row == 0 {
				return visited.Cardinality()
			}
		} else if dir == 'R' {
			for i := position.col; i < len(rows[0]); i++ {
				if rows[position.row][i] == '#' {
					dir = newDirection(dir)
					break
				} else {
					position.col = i
					visited.Add(position)
				}
			}
			if position.col == len(rows[0])-1 {
				return visited.Cardinality()
			}
		} else if dir == 'B' {
			for i := position.row; i < len(rows); i++ {
				if rows[i][position.col] == '#' {
					dir = newDirection(dir)
					break
				} else {
					position.row = i
					visited.Add(position)
				}
			}
			if position.row == len(rows)-1 {
				return visited.Cardinality()
			}
		} else if dir == 'L' {
			for i := position.col; i >= 0; i-- {
				if rows[position.row][i] == '#' {
					dir = newDirection(dir)
					break
				} else {
					position.col = i
					visited.Add(position)
				}
			}
			if position.col == 0 {
				return visited.Cardinality()
			}
		}
	}
}

func part2(input string) int {
	strings.Split(input, "\n")
	return 0
}

func main() {
	input := common.ReadInput(6)
	fmt.Println("DAY 06")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
