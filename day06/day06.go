package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/paolostyle/advent-of-code-2024/common"
)

type pos struct {
	col int
	row int
}

type node struct {
	position  pos
	direction rune
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

func move(dir rune, position pos) pos {
	switch dir {
	case 'T':
		return pos{col: position.col, row: position.row - 1}
	case 'R':
		return pos{col: position.col + 1, row: position.row}
	case 'B':
		return pos{col: position.col, row: position.row + 1}
	case 'L':
		return pos{col: position.col - 1, row: position.row}
	}
	return pos{}
}

func checkBoundary(dir rune, position pos, rows []string) bool {
	switch dir {
	case 'T':
		return position.row == 0
	case 'R':
		return position.col == len(rows[0])-1
	case 'B':
		return position.row == len(rows)-1
	case 'L':
		return position.col == 0
	}
	return false
}

func getInitialPosition(rows []string) (pos, error) {
	for i, row := range rows {
		col := strings.Index(row, "^")
		if col != -1 {
			return pos{col: col, row: i}, nil
		}
	}

	return pos{}, errors.New("no initial position found")
}

func run(rows []string, position pos) (mapset.Set[pos], error) {
	visited := mapset.NewSet[node]()
	visited.Add(node{position: position, direction: 'T'})

	dir := 'T'
	for {
		nextMove := move(dir, position)
		if rows[nextMove.row][nextMove.col] == '#' {
			dir = newDirection(dir)
		} else {
			position = nextMove

			if visited.Contains(node{position: position, direction: dir}) {
				return nil, errors.New("loop detected")
			}

			visited.Add(node{position: position, direction: dir})
		}

		if checkBoundary(dir, position, rows) {
			uniquePositions := mapset.NewSet[pos]()
			for _, node := range visited.ToSlice() {
				uniquePositions.Add(node.position)
			}
			return uniquePositions, nil
		}
	}
}

func part1(input string) int {
	rows := strings.Split(input, "\n")
	position, err := getInitialPosition(rows)
	if err != nil {
		log.Fatal(err)
	}

	visited, err := run(rows, position)
	if err != nil {
		log.Fatal("no loops expected in initial run")
	}

	return visited.Cardinality()
}

func part2(input string) int {
	rows := strings.Split(input, "\n")
	initialPos, err := getInitialPosition(rows)
	if err != nil {
		log.Fatal(err)
	}

	initialVisited, err := run(rows, initialPos)
	if err != nil {
		log.Fatal("no loops expected in initial run")
	}

	loops := 0

	for _, pos := range initialVisited.ToSlice() {
		if pos == initialPos {
			continue
		}

		rowsCopy := make([]string, len(rows))
		copy(rowsCopy, rows)
		row := []rune(rows[pos.row])
		row[pos.col] = '#'
		rowsCopy[pos.row] = string(row)

		_, err := run(rowsCopy, initialPos)
		if err != nil {
			loops += 1
		}
	}

	return loops
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(6)
	fmt.Println("DAY 06")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
