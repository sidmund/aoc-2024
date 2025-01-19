package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sidmund/aoc-2024/lib"
)

func observe(stones map[int]int, blinks int) int {
	defer lib.Measure(time.Now(), "Observe")

	for range blinks {
		observed := map[int]int{}
		for stone, amount := range stones {
			if stone == 0 {
				observed[1] += amount
			} else if s := strconv.Itoa(stone); len(s)%2 == 0 {
				left, _ := strconv.Atoi(s[:len(s)/2])
				right, _ := strconv.Atoi(s[len(s)/2:])
				observed[left] += amount
				observed[right] += amount
			} else {
				observed[stone*2024] += amount
			}
		}
		stones = observed
	}

	total := 0
	for _, amount := range stones {
		total += amount
	}
	return total
}

func main() {
	raw, _ := os.ReadFile("day11/input")
	stones := map[int]int{}
	for _, s := range strings.Fields(string(raw)) {
		i, _ := strconv.Atoi(s)
		stones[i]++
	}

	for i, blinks := range []int{25, 75} {
		fmt.Printf("Part %d: %d\n", i+1, observe(stones, blinks))
	}
}
