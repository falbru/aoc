package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func partOne(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	distance_sum := 0
	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		if distance < 0 {
			distance = -distance
		}

		distance_sum += distance
	}

	return distance_sum
}

func partTwo(left []int, right []int) int {
	right_num_count := make(map[int]int)

	for _, num := range right {
		count, ok := right_num_count[num]
		if ok {
			right_num_count[num] = count + 1
		} else {
			right_num_count[num] = 1
		}
	}

	similarity_score := 0
	for _, num := range left {
		count, ok := right_num_count[num]
		if ok {
			similarity_score += num * count
		}
	}

	return similarity_score
}

func main() {
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		pairs := strings.Split(lines[i], "   ")

		if len(pairs) != 2 {
			continue
		}

		left[i], _ = strconv.Atoi(pairs[0])
		right[i], _ = strconv.Atoi(pairs[1])
	}

	ans1 := partOne(left, right)
	ans2 := partTwo(left, right)

	fmt.Println("PART ONE ANSWER:", ans1)
	fmt.Println("PART TWO ANSWER:", ans2)
}
