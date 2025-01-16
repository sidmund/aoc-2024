package main

import (
	"fmt"
	"slices"

	"github.com/sidmund/aoc-2024/lib"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	lines := lib.ReadLines("day01/input")

	var left, right []int
	frequency := map[int]int{}
	for _, line := range lines {
		var l, r int
		fmt.Sscanf(line, "%d   %d", &l, &r)
		left, right = append(left, l), append(right, r)
		frequency[r]++
	}

	slices.Sort(left)
	slices.Sort(right)

	sum, similarityScore := 0, 0
	for i := range left {
		sum += abs(left[i] - right[i])
		similarityScore += frequency[left[i]] * left[i]
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", similarityScore)
}
