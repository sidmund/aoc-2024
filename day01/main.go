package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	lib "github.com/sidmund/aoc-2024/lib"
)

func abs(n int64) int64 {
    if n < 0 {
        return -n
    }
    return n
}

func main() {
	lines := lib.ReadLines("day01/input")

	left := make([]int64, len(lines))
	right := make([]int64, len(lines))
	for _, line := range lines {
		split := strings.Fields(line)
		l, _ := strconv.ParseInt(split[0], 10, 64)
		r, _ := strconv.ParseInt(split[1], 10, 64)
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	var sum int64
	for i := range left {
		sum += abs(left[i] - right[i])
	}
    fmt.Println("Part 1:", sum)

    var similarityScore int64
    for _, l := range left {
        var frequency int64
        for _, r := range right {
            if l == r {
                frequency++
            }
        }
        similarityScore += frequency * l
    }
    fmt.Println("Part 2:", similarityScore)
}
