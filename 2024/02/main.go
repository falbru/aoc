package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func isSafeValue(left int, right int, ascending bool) bool {
	if (ascending && left > right) || (!ascending && left < right) {
		return false
	}

	diff := right - left
	if diff < 0 {
		diff = -diff
	}

	if diff < 1 || diff > 3 {
		return false
	}

	return true
}

func isSafeReport(levels []int) (bool, int) {
	ascendingNumbers := 0
	for i := range len(levels) - 1 {
		if levels[i+1] > levels[i] {
			ascendingNumbers++
		}
	}

	ascending := false
	if ascendingNumbers > len(levels)/2 {
		ascending = true
	}

	for i := range len(levels) - 1 {
		if !isSafeValue(levels[i], levels[i+1], ascending) {
			return false, i
		}
	}
	return true, -1
}

func partOne(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		isSafe, _ := isSafeReport(report)
		if isSafe {
			safeReports += 1
		}
	}

	return safeReports
}

func partTwo(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		isSafe, unsafeValue := isSafeReport(report)

		if !isSafe {
			updatedReport1 := lo.Filter(report, func(item int, index int) bool {
				return index != unsafeValue
			})
			updatedReport2 := lo.Filter(report, func(item int, index int) bool {
				return index != unsafeValue+1
			})

			isSafe1, _ := isSafeReport(updatedReport1)
			isSafe2, _ := isSafeReport(updatedReport2)

			if isSafe1 || isSafe2 {
				isSafe = true
			}
		}

		if isSafe {
			safeReports++
		}
	}

	return safeReports
}

func main() {
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	nonEmptyLines := lo.Filter(lines, func(line string, index int) bool {
		return line != ""
	})

	splitLines := lo.Map(nonEmptyLines, func(line string, index int) []string {
		return strings.Split(line, " ")
	})

	reports := lo.Map(splitLines, func(x []string, index int) []int {
		return lo.Map(x, func(y string, index int) int {
			result, _ := strconv.Atoi(y)
			return result
		})
	})

	fmt.Println("PART ONE ANSWER:", partOne(reports))
	fmt.Println("PART TWO ANSWER:", partTwo(reports))
}
