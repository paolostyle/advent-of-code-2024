package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

type coord struct {
	x int
	y int
}

func getTopoMap(input string) [][]int {
	rows := strings.Split(input, "\n")
	topoMap := make([][]int, len(rows))
	for i, row := range rows {
		topoMap[i] = common.StringsToNumberSafe(strings.Split(row, ""), -1)
	}
	return topoMap
}

func getAdjacencyMapAndStartPos(topoMap [][]int) (map[coord][]coord, []coord) {
	adjMap := make(map[coord][]coord)
	startingPoints := []coord{}

	for y, row := range topoMap {
		for x, val := range row {
			adjMap[coord{x, y}] = []coord{}

			if val == -1 {
				continue
			}

			if val == 0 {
				startingPoints = append(startingPoints, coord{x, y})
			}

			// left
			if x > 0 && topoMap[y][x-1] == val+1 {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x - 1, y})
			}

			// right
			if x < len(row)-1 && topoMap[y][x+1] == val+1 {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x + 1, y})
			}

			// up
			if y > 0 && topoMap[y-1][x] == val+1 {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x, y - 1})
			}

			// down
			if y < len(topoMap)-1 && topoMap[y+1][x] == val+1 {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x, y + 1})
			}
		}
	}

	return adjMap, startingPoints
}

func traverse(adjMap map[coord][]coord, start coord, visited map[coord]bool, path []coord) []coord {
	visited[start] = true
	path = append(path, start)

	for _, neighbor := range adjMap[start] {
		if !visited[neighbor] {
			path = traverse(adjMap, neighbor, visited, path)
		}
	}

	return path
}

func findPeaks(adjMap map[coord][]coord, topoMap [][]int, start coord) []coord {
	path := traverse(adjMap, start, make(map[coord]bool), []coord{})
	peaks := []coord{}

	for _, p := range path {
		if topoMap[p.y][p.x] == 9 {
			peaks = append(peaks, p)
		}
	}

	return peaks
}

func findPathsToPeak(adjMap map[coord][]coord, start coord, end coord, visited map[coord]bool, path []coord, pathsCount *int) {
	visited[start] = true

	if start == end {
		*pathsCount += 1
	} else {
		for _, neighbor := range adjMap[start] {
			if !visited[neighbor] {
				path = append(path, start)
				findPathsToPeak(adjMap, neighbor, end, visited, path, pathsCount)
				path = path[:len(path)-1]
			}
		}
	}

	visited[start] = false
}

func part1(input string) int {
	topoMap := getTopoMap(input)
	adjMap, startingPoints := getAdjacencyMapAndStartPos(topoMap)

	score := 0
	for _, start := range startingPoints {
		score += len(findPeaks(adjMap, topoMap, start))
	}

	return score
}

func part2(input string) int {
	topoMap := getTopoMap(input)
	adjMap, startingPoints := getAdjacencyMapAndStartPos(topoMap)

	totalRating := 0
	for _, start := range startingPoints {
		rating := 0
		peaks := findPeaks(adjMap, topoMap, start)

		for _, p := range peaks {
			findPathsToPeak(adjMap, start, p, make(map[coord]bool), []coord{}, &rating)
		}

		totalRating += rating
	}

	return totalRating
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(10)
	fmt.Println("DAY 10")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
