package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sidmund/aoc-2024/lib"
)

func process(lines []string) int {
	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, line := range lines {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum += a * b
		}
	}
	return sum
}

func processWithConditionals(lines []string) int {
	sum := 0
	enabled := true
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	for _, line := range lines {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			if match[0][2] == '(' {
				enabled = true
			} else if match[0][2] == 'n' {
				enabled = false
			} else if enabled && match[0][2] == 'l' {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				sum += a * b
			}
		}
	}
	return sum
}

func main() {
	lines := lib.ReadLines("day03/input")

	fmt.Println("Part 1:", process(lines))
	fmt.Println("Part 2:", processWithConditionals(lines))
}
