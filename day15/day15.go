package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

type coord struct {
	x, y int
}

func nextPosition(dir rune, pos coord) coord {
	switch dir {
	case '^':
		return coord{pos.x, pos.y - 1}
	case '>':
		return coord{pos.x + 1, pos.y}
	case 'v':
		return coord{pos.x, pos.y + 1}
	case '<':
		return coord{pos.x - 1, pos.y}
	}
	return coord{}
}

func swap(field []string, oldPos coord, newPos coord) {
	tmp := field[oldPos.y][oldPos.x]
	field[oldPos.y] = field[oldPos.y][:oldPos.x] + string(field[newPos.y][newPos.x]) + field[oldPos.y][oldPos.x+1:]
	field[newPos.y] = field[newPos.y][:newPos.x] + string(tmp) + field[newPos.y][newPos.x+1:]
}

func move(field []string, pos coord, dir rune) coord {
	if pos.x < 0 || pos.y < 0 || pos.y >= len(field) || pos.x >= len(field[pos.y]) {
		return pos
	}
	nextPos := nextPosition(dir, pos)
	nextVal := field[nextPos.y][nextPos.x]

	if nextVal == '#' {
		return pos
	}

	if nextVal == '.' {
		swap(field, pos, nextPos)
		return nextPos
	}

	if nextVal == 'O' {
		moved := move(field, nextPos, dir)
		if moved != nextPos {
			swap(field, pos, nextPos)
			return nextPos
		}
	}

	if nextVal == '[' || nextVal == ']' {

	}

	return pos
}

func dbg(field []string) {
	for _, line := range field {
		fmt.Println(line)
	}
	fmt.Println()
}

func part1(input string) int {
	splitInput := strings.Split(input, "\n\n")
	field := strings.Split(splitInput[0], "\n")
	instructions := strings.Join(strings.Split(splitInput[1], "\n"), "")

	pos := coord{0, 0}

	for y, line := range field {
		if x := strings.Index(line, "@"); x != -1 {
			pos = coord{x, y}
			break
		}
	}

	for _, dir := range instructions {
		pos = move(field, pos, dir)
	}

	gpsSum := 0

	for y, line := range field {
		for x, val := range line {
			if val == 'O' {
				gpsSum += y*100 + x
			}
		}
	}

	return gpsSum
}

func newMap(field []string) []string {
	newField := make([]string, len(field))
	for i, line := range field {
		var sb strings.Builder

		for _, val := range line {
			if val == 'O' {
				sb.WriteRune('[')
				sb.WriteRune(']')
			} else if val == '@' {
				sb.WriteRune('@')
				sb.WriteRune('.')
			} else {
				sb.WriteRune(val)
				sb.WriteRune(val)
			}
		}

		newField[i] = sb.String()
	}

	dbg(newField)

	return newField
}

func part2(input string) int {
	splitInput := strings.Split(input, "\n\n")
	field := strings.Split(splitInput[0], "\n")
	// instructions := strings.Join(strings.Split(splitInput[1], "\n"), "")

	newMap(field)
	return 0
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(15)
	fmt.Println("DAY 15")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
