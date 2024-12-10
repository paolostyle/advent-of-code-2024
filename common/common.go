package common

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadInput(day int) string {
	isTest := flag.Bool("test", false, "Use test data")
	flag.Parse()

	directory := fmt.Sprintf("./day%02d/", day)
	filename := directory + "input.txt"
	if *isTest {
		filename = directory + "test_input.txt"
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(data))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func StringsToNumbers(strings []string) []int {
	nums := make([]int, len(strings))
	for i, str := range strings {
		nums[i] = StrToInt(str)
	}
	return nums
}

func StringsToNumberSafe(strings []string, defaultValue int) []int {
	nums := make([]int, len(strings))
	for i, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			nums[i] = defaultValue
		} else {
			nums[i] = num
		}
	}
	return nums
}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("took", elapsed)
}
