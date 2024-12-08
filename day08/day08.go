package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/paolostyle/advent-of-code-2024/common"
)

type coord struct {
	x int
	y int
}

func getVector(a coord, b coord) coord {
	return coord{x: a.x - b.x, y: a.y - b.y}
}

func getAntinode(a coord, b coord) coord {
	vec := getVector(a, b)
	return coord{x: a.x + vec.x, y: a.y + vec.y}
}

func isInBounds(c coord, size int) bool {
	return (c.x >= 0 && c.x < size) && (c.y >= 0 && c.y < size)
}

func getAntennas(field []string) map[string][]coord {
	antennaRegex := regexp.MustCompile("[0-9a-zA-Z]{1}")
	antennas := make(map[string][]coord)

	for y, row := range field {
		matches := antennaRegex.FindAllStringIndex(row, -1)
		for _, match := range matches {
			x := match[0]
			antennaType := string(row[x])
			antennas[antennaType] = append(antennas[antennaType], coord{x, y})
		}
	}

	return antennas
}

func part1(input string) int {
	field := strings.Split(input, "\n")
	antennas := getAntennas(field)
	size := len(field)
	antinodes := mapset.NewSet[coord]()

	for _, coords := range antennas {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				antinodeA := getAntinode(coords[i], coords[j])
				if isInBounds(antinodeA, size) {
					antinodes.Add(antinodeA)
				}

				antinodeB := getAntinode(coords[j], coords[i])
				if isInBounds(antinodeB, size) {
					antinodes.Add(antinodeB)
				}
			}
		}
	}

	return antinodes.Cardinality()
}

func part2(input string) int {
	field := strings.Split(input, "\n")
	antennas := getAntennas(field)
	size := len(field)
	antinodes := mapset.NewSet[coord]()

	for _, coords := range antennas {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				antinodes.Add(coords[i])
				antinodes.Add(coords[j])

				vec := getVector(coords[i], coords[j])
				antinode := coords[i]
				for {
					antinode = coord{x: antinode.x + vec.x, y: antinode.y + vec.y}
					if isInBounds(antinode, size) {
						antinodes.Add(antinode)
					} else {
						break
					}
				}

				antinode = coords[j]
				for {
					antinode = coord{x: antinode.x - vec.x, y: antinode.y - vec.y}
					if isInBounds(antinode, size) {
						antinodes.Add(antinode)
					} else {
						break
					}
				}
			}
		}
	}

	return antinodes.Cardinality()
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(8)
	fmt.Println("DAY 08")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
