package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/samber/lo"
)

func partOne(lines []string) int {
	N := len(lines)

	horizontal := lines
	vertical := make([]string, N)
	right_diagonal := make([]string, N*2-1)
	left_diagonal := make([]string, N*2-1)

	for i := range N {
		for j := range N {
			vertical[i] += string(lines[j][i])
		}
	}

	for i := range N {
		for j := range i + 1 {
			right_diagonal[i] += string(lines[j][N-1-i+j])
			if i < N-1 {
				right_diagonal[2*N-2-i] += string(lines[N-1-i+j][j])
			}
		}
	}

	for i := range N {
		for j := range i + 1 {
			left_diagonal[i] += string(lines[j][i-j])
			if i < N-1 {
				left_diagonal[2*N-2-i] += string(lines[N-1-i+j][N-1-j])
			}
		}
	}

	total_lines := slices.Concat(horizontal, vertical, left_diagonal, right_diagonal)

	count := 0
	for _, line := range total_lines {
		count += strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
	}

	return count
}

func partTwo(lines []string) int {
	N := len(lines)

	count := 0

	for i := 1; i < N-1; i++ {
		for j := 1; j < N-1; j++ {
			diag1 := string(lines[i-1][j-1]) + string(lines[i][j]) + string(lines[i+1][j+1])
			diag2 := string(lines[i+1][j-1]) + string(lines[i][j]) + string(lines[i-1][j+1])

			if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
				count++
			}
		}
	}

	return count
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := lo.Filter(strings.Split(string(input), "\n"), func(line string, index int) bool {
		return line != ""
	})

	fmt.Println("PART ONE ANSWER:", partOne(lines))
	fmt.Println("PART TWO ANSWER:", partTwo(lines))
}
