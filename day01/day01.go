package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

func splitIntoLists(input string) ([]int, []int) {
	locations := strings.Split(input, "\n")
	leftList := []int{}
	rightList := []int{}

	for _, location := range locations {
		parts := strings.Fields(location)

		leftLocation, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		rightLocation, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		leftList = append(leftList, leftLocation)
		rightList = append(rightList, rightLocation)
	}

	return leftList, rightList
}

func part1(input string) int {
	leftList, rightList := splitIntoLists(input)

	slices.Sort(leftList)
	slices.Sort(rightList)

	sum := 0
	for i, left := range leftList {
		right := rightList[i]
		sum += common.Abs(right - left)
	}

	return sum
}

func part2(input string) int {
	leftList, rightList := splitIntoLists(input)
	countList := make(map[int]int)

	for _, right := range rightList {
		countList[right] += 1
	}

	similarity := 0
	for _, left := range leftList {
		similarity += left * countList[left]
	}

	return similarity
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(1)
	fmt.Println("DAY 01")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
