package main

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/paolostyle/advent-of-code-2024/common"
)

type rule struct {
	after  mapset.Set[string]
	before mapset.Set[string]
}

func newRule() *rule {
	r := rule{
		after:  mapset.NewSet[string](),
		before: mapset.NewSet[string](),
	}

	return &r
}

func prepareRuleset(input string) ([]string, map[string]*rule) {
	split := strings.Split(input, "\n\n")
	rules := strings.Split(split[0], "\n")
	updates := strings.Split(split[1], "\n")

	ruleset := make(map[string]*rule)
	for _, rule := range rules {
		pages := strings.Split(rule, "|")

		_, exists := ruleset[pages[0]]

		if !exists {
			ruleset[pages[0]] = newRule()
		}

		ruleset[pages[0]].after.Add(pages[1])

		_, exists = ruleset[pages[1]]

		if !exists {
			ruleset[pages[1]] = newRule()
		}

		ruleset[pages[1]].before.Add(pages[0])
	}

	return updates, ruleset
}

func splitUpdates(updates []string, ruleset map[string]*rule) ([][]string, [][]string) {
	valid := make([][]string, 0)
	invalid := make([][]string, 0)

	for _, updateList := range updates {
		update := strings.Split(updateList, ",")
		isValid := true

	checkLoop:
		for i, page := range update {
			pagesBefore := update[:i]
			pagesAfter := update[i+1:]

			for _, pageBefore := range pagesBefore {
				if ruleset[page].after.Contains(pageBefore) {
					isValid = false
					break checkLoop
				}
			}

			for _, pageAfter := range pagesAfter {
				if ruleset[page].before.Contains(pageAfter) {
					isValid = false
					break checkLoop
				}
			}
		}

		if isValid {
			valid = append(valid, update)
		} else {
			invalid = append(invalid, update)
		}
	}

	return valid, invalid
}

func getMiddlePage(update []string) int {
	return common.StrToInt(update[int(len(update)/2)])
}

func part1(input string) int {
	updates, ruleset := prepareRuleset(input)
	valid, _ := splitUpdates(updates, ruleset)

	pagesSum := 0

	for _, update := range valid {
		pagesSum += getMiddlePage(update)
	}

	return pagesSum
}

func part2(input string) int {
	updates, ruleset := prepareRuleset(input)
	_, invalid := splitUpdates(updates, ruleset)

	pagesSum := 0

	for _, update := range invalid {
	rewind:
		for i, page := range update {
			pagesBefore := update[:i]
			pagesAfter := update[i+1:]

			for j, pageBefore := range pagesBefore {
				if ruleset[page].after.Contains(pageBefore) {
					update[i], update[j] = update[j], update[i]
					goto rewind
				}
			}

			for j, pageAfter := range pagesAfter {
				if ruleset[page].before.Contains(pageAfter) {
					update[i], update[j+i+1] = update[j+i+1], update[i]
					goto rewind
				}
			}
		}

		pagesSum += getMiddlePage(update)
	}

	return pagesSum

}

func main() {
	input := common.ReadInput(5)
	fmt.Println("DAY 05")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
