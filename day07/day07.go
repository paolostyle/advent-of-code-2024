package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/paolostyle/advent-of-code-2024/common"
)

func parseOperation(operation string) (int, []int) {
	split := strings.Split(operation, ":")
	result := common.StrToInt(split[0])
	partsStr := strings.Split(strings.TrimSpace(split[1]), " ")
	parts := make([]int, 0)
	for _, part := range partsStr {
		parts = append(parts, common.StrToInt(part))
	}
	return result, parts
}

func getOperators(n int, withConcat bool) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	subCombinations := getOperators(n-1, withConcat)
	combinations := [][]string{}

	for _, subCombination := range subCombinations {
		combinations = append(combinations, append([]string{"*"}, subCombination...))
		combinations = append(combinations, append([]string{"+"}, subCombination...))
		if withConcat {
			combinations = append(combinations, append([]string{"||"}, subCombination...))
		}
	}

	return combinations
}

func runOperations(operations []string, withConcat bool) int {
	sum := 0
	for _, operation := range operations {
		result, parts := parseOperation(operation)
		operatorsList := getOperators(len(parts)-1, withConcat)

		for _, operators := range operatorsList {
			calcResult := parts[0]
			for i := 0; i < len(operators); i++ {
				if operators[i] == "+" {
					calcResult += parts[i+1]
				} else if operators[i] == "*" {
					calcResult *= parts[i+1]
				} else if operators[i] == "||" {
					calcResult = common.StrToInt(strconv.Itoa(calcResult) + strconv.Itoa(parts[i+1]))
				}
			}
			if calcResult == result {
				sum += result
				break
			}
		}
	}
	return sum
}

func part1(input string) int {
	operations := strings.Split(input, "\n")
	return runOperations(operations, false)
}

func part2(input string) int {
	operations := strings.Split(input, "\n")
	return runOperations(operations, true)
}

func main() {
	defer common.TimeTrack(time.Now())
	input := common.ReadInput(7)
	fmt.Println("DAY 07")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
