package main

import (
	"fmt"
	"maps"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

type coord struct {
	x int
	y int
}

func getAdjacencyMap(field []string) map[coord][]coord {
	adjMap := make(map[coord][]coord)

	for y, row := range field {
		for x, val := range []byte(row) {
			adjMap[coord{x, y}] = []coord{}

			// left
			if x > 0 && field[y][x-1] == val {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x - 1, y})
			}

			// right
			if x < len(row)-1 && field[y][x+1] == val {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x + 1, y})
			}

			// up
			if y > 0 && field[y-1][x] == val {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x, y - 1})
			}

			// down
			if y < len(field)-1 && field[y+1][x] == val {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x, y + 1})
			}
		}
	}

	return adjMap
}

func differentNeighbors(field []string, pos coord) int {
	value := field[pos.y][pos.x]

	topIsDiff := pos.y == 0 || (pos.y > 0 && field[pos.y-1][pos.x] != value)
	bottomIsDiff := pos.y == len(field)-1 || (pos.y < len(field)-1 && field[pos.y+1][pos.x] != value)
	leftIsDiff := pos.x == 0 || (pos.x > 0 && field[pos.y][pos.x-1] != value)
	rightIsDiff := pos.x == len(field[0])-1 || (pos.x < len(field[0])-1 && field[pos.y][pos.x+1] != value)

	return common.BoolToInt(topIsDiff) + common.BoolToInt(bottomIsDiff) + common.BoolToInt(leftIsDiff) + common.BoolToInt(rightIsDiff)
}

func traverse(adjMap map[coord][]coord, field []string, pos coord, visited map[coord]bool, perimeter *int) map[coord]bool {
	visited[pos] = true
	*perimeter += differentNeighbors(field, pos)

	for _, neighbor := range adjMap[pos] {
		if !visited[neighbor] {
			traverse(adjMap, field, neighbor, visited, perimeter)
		}
	}

	return visited
}

func part1(input string) int {
	field := strings.Split(input, "\n")
	adjMap := getAdjacencyMap(field)
	totalVisited := make(map[coord]bool)
	price := 0

	for i, row := range field {
		for j := range row {
			if !totalVisited[coord{j, i}] {
				perimeter := 0
				visited := traverse(adjMap, field, coord{j, i}, make(map[coord]bool), &perimeter)
				maps.Copy(totalVisited, visited)
				price += perimeter * len(visited)
			}
		}
	}

	return price
}

func part2(input string) int {
	return 0
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(12)
	fmt.Println("DAY 12")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
