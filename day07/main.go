package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sidmund/aoc-2024/lib"
)

// Shrink numbers down to a single value by trying out all possible combinations of operators
func calibrate(target int, ns []int, concatenate bool) int {
	if len(ns) == 1 {
		if ns[0] == target {
			return target
		}
		return 0
	}

	if result := calibrate(target, append([]int{ns[0] + ns[1]}, ns[2:]...), concatenate); result > 0 {
		return result
	}
	if result := calibrate(target, append([]int{ns[0] * ns[1]}, ns[2:]...), concatenate); result > 0 {
		return result
	}
	if concatenate {
		concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", ns[0], ns[1]))
		return calibrate(target, append([]int{concatenated}, ns[2:]...), concatenate)
	}

	return 0
}

func main() {
	lines := lib.ReadLines("day07/input")

	defer lib.Measure(time.Now(), "Day 7")
	part1, part2 := 0, 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		operands := []int{}
		for _, value := range strings.Fields(parts[1]) {
			operand, _ := strconv.Atoi(value)
			operands = append(operands, operand)
		}

		part1 += calibrate(target, operands, false)
		part2 += calibrate(target, operands, true)
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}
