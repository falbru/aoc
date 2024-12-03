package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func partOne(input string) int {
	r, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")
	muls := r.FindAllString(input, -1)

	sum := 0
	for _, mul := range muls {
		factors := strings.Split(mul[4:len(mul)-1], ",")

		a, _ := strconv.Atoi(factors[0])
		b, _ := strconv.Atoi(factors[1])
		sum += a * b
	}

	return sum
}

func partTwo(input string) int {
	r, _ := regexp.Compile("(mul\\(\\d+,\\d+\\))|(don't\\(\\))|(do\\(\\))")
	insts := r.FindAllString(input, -1)

	do := true
	sum := 0

	for _, inst := range insts {
		if inst == "do()" {
			do = true
		} else if inst == "don't()" {
			do = false
		} else if do && inst[0:3] == "mul" {
			factors := strings.Split(inst[4:len(inst)-1], ",")

			a, _ := strconv.Atoi(factors[0])
			b, _ := strconv.Atoi(factors[1])
			sum += a * b
		}
	}

	return sum
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("PART ONE ANSWER:", partOne(string(input)))
	fmt.Println("PART TWO ANSWER:", partTwo(string(input)))
}
