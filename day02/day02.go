package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/paolostyle/advent-of-code-2024/common"
)

func getSafeReportsCount(input string, allowError bool) int {
	reportLines := strings.Split(input, "\n")
	safeReports := 0

	for _, reportLine := range reportLines {
		report := common.StringsToNumbers(strings.Fields(reportLine))
		isSafe := isReportSafe(report, allowError)

		if isSafe {
			safeReports += 1
		}
	}

	return safeReports
}

func isReportSafe(report []int, allowError bool) bool {
	shouldIncrease := report[1] > report[0]
	shouldDecrease := report[1] < report[0]

	isSafe := true
	for i, num := range report {
		if i == 0 {
			continue
		}
		prev := report[i-1]
		diff := common.Abs(num - prev)

		isWrongDirection := shouldIncrease && num < prev || shouldDecrease && num > prev
		isDiffInRange := diff >= 1 && diff <= 3

		if isWrongDirection || !isDiffInRange {
			isSafe = false
			break
		}
	}

	if !isSafe && allowError {
		for i := range report {
			dampenedReport := slices.Clone(report[:i])
			dampenedReport = append(dampenedReport, report[i+1:]...)
			isSafe = isReportSafe(dampenedReport, false)

			if isSafe {
				break
			}
		}
	}

	return isSafe
}

func part1(input string) int {
	return getSafeReportsCount(input, false)
}

func part2(input string) int {
	return getSafeReportsCount(input, true)
}

func main() {
	input := common.ReadInput(2)
	fmt.Println("DAY 02")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
