package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func inValidOrder(left int, right int, deps map[int]map[int]bool) bool {
	return deps[left][right] || !deps[right][left]
}

func partOne(updates [][]int, deps map[int]map[int]bool) int {
	sum := 0

	for _, nums := range updates {
		correctOrder := true

		for i := 1; i < len(nums); i++ {
			if !inValidOrder(nums[i-1], nums[i], deps) {
				correctOrder = false
			}
		}

		if correctOrder {
			sum += nums[len(nums)/2]
		}
	}

	return sum
}

func partTwo(updates [][]int, deps map[int]map[int]bool) int {
	sum := 0

	for _, nums := range updates {
		correctOrder := true

		for i := 0; i < len(nums); i++ {
			for j := i; j > 0 && !inValidOrder(nums[j-1], nums[j], deps); j-- {
				correctOrder = false
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}

		if !correctOrder {
			sum += nums[len(nums)/2]
		}
	}

	return sum
}

func main() {
	input, _ := os.ReadFile("input.txt")

	sections := strings.Split(string(input), "\n\n")

	var updates [][]int
	dependencies := make(map[int]map[int]bool)

	for _, line := range strings.Split(sections[0], "\n") {
		nums := strings.Split(line, "|")

		left, _ := strconv.Atoi(nums[0])
		right, _ := strconv.Atoi(nums[1])

		if _, exists := dependencies[left]; !exists {
			dependencies[left] = make(map[int]bool)
		}
		dependencies[left][right] = true
	}

	for _, line := range strings.Split(sections[1], "\n") {
		update := lo.Map(strings.Split(line, ","), func(item string, index int) int {
			result, _ := strconv.Atoi(item)
			return result
		})

		updates = append(updates, update)
	}

	fmt.Println("PART ONE ANSWER:", partOne(updates, dependencies))
	fmt.Println("PART TWO ANSWER:", partTwo(updates, dependencies))
}
