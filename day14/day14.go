package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

type coord struct {
	x, y int
}

type point struct {
	position coord
	velocity coord
}

var lineRegex = regexp.MustCompile(`p=([-\d]+),([-\d]+) v=([-\d]+),([-\d]+)`)

func mod(a, b int) int {
	return (a%b + b) % b
}

const xSize = 101
const ySize = 103

func part1(input string) int {
	lines := strings.Split(input, "\n")

	verticalSep := xSize / 2
	horizontalSep := ySize / 2
	iterations := 100

	counts := make([]int, 4)

	for _, line := range lines {
		matches := lineRegex.FindAllStringSubmatch(line, -1)[0]
		position := coord{common.StrToInt(matches[1]), common.StrToInt(matches[2])}
		velocity := coord{common.StrToInt(matches[3]), common.StrToInt(matches[4])}

		position.x = mod(position.x+velocity.x*iterations, xSize)
		position.y = mod(position.y+velocity.y*iterations, ySize)

		if position.x < verticalSep && position.y < horizontalSep {
			counts[0]++
		}
		if position.x > verticalSep && position.y < horizontalSep {
			counts[1]++
		}
		if position.x < verticalSep && position.y > horizontalSep {
			counts[2]++
		}
		if position.x > verticalSep && position.y > horizontalSep {
			counts[3]++
		}
	}

	return counts[0] * counts[1] * counts[2] * counts[3]
}

func getAdjacencyMap(field [ySize][xSize]bool) map[coord][]coord {
	adjMap := make(map[coord][]coord)

	for y, row := range field {
		for x, val := range row {
			adjMap[coord{x, y}] = []coord{}

			if !val {
				continue
			}

			// left
			if x > 0 && field[y][x-1] {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x - 1, y})
			}

			// right
			if x < len(row)-1 && field[y][x+1] {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x + 1, y})
			}

			// up
			if y > 0 && field[y-1][x] {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x, y - 1})
			}

			// down
			if y < len(field)-1 && field[y+1][x] {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x, y + 1})
			}

			// top left
			if x > 0 && y > 0 && field[y-1][x-1] {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x - 1, y - 1})
			}

			// top right
			if x < len(row)-1 && y > 0 && field[y-1][x+1] {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x + 1, y - 1})
			}

			// bottom left
			if x > 0 && y < len(field)-1 && field[y+1][x-1] {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x - 1, y + 1})
			}

			// bottom right
			if x < len(row)-1 && y < len(field)-1 && field[y+1][x+1] {
				adjMap[coord{x, y}] = append(adjMap[coord{x, y}], coord{x + 1, y + 1})
			}
		}
	}

	return adjMap
}

func dfs(adjMap map[coord][]coord, start coord, visited map[coord]bool, path int) int {
	visited[start] = true
	maxPath := path

	for _, neighbor := range adjMap[start] {
		if !visited[neighbor] {
			currentPath := dfs(adjMap, neighbor, visited, path+1)
			if currentPath > maxPath {
				maxPath = currentPath
			}
		}
	}

	return maxPath
}

// this solution is terrible but idc at this point
func part2(input string) int {
	lines := strings.Split(input, "\n")

	area := [ySize][xSize]bool{}
	points := make([]point, len(lines))

	for _, line := range lines {
		matches := lineRegex.FindAllStringSubmatch(line, -1)[0]
		position := coord{common.StrToInt(matches[1]), common.StrToInt(matches[2])}
		velocity := coord{common.StrToInt(matches[3]), common.StrToInt(matches[4])}
		points = append(points, point{position, velocity})
	}

	globalLongestPath := 0
	longestPathIteration := 0

	for i := 0; i < ySize*xSize; i++ {
		for j, point := range points {
			points[j].position.x = mod(point.position.x+point.velocity.x, xSize)
			points[j].position.y = mod(point.position.y+point.velocity.y, ySize)

			area[point.position.y][point.position.x] = true
		}

		adjMap := getAdjacencyMap(area)
		visited := make(map[coord]bool)
		maxPath := 0
		for node := range adjMap {
			pathLength := dfs(adjMap, node, visited, 1)
			if pathLength > maxPath {
				maxPath = pathLength
			}
		}

		if maxPath > globalLongestPath {
			globalLongestPath = maxPath
			longestPathIteration = i
		}

		area = [ySize][xSize]bool{}
	}

	return longestPathIteration
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(14)
	fmt.Println("DAY 14")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
