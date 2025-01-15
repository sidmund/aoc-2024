package main

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"

    "github.com/sidmund/aoc-2024/lib"
)

func mul(ins string) int {
    ops := strings.Split(ins[4:len(ins)-1], ",")
    a, _ := strconv.Atoi(ops[0])
    b, _ := strconv.Atoi(ops[1])
    return a*b
}

func process(lines []string) int {
    sum := 0
    re := regexp.MustCompile(`mul\(\d+,\d+\)`)
    for _, line := range lines {
        for _, match := range re.FindAllString(line, -1) {
            sum += mul(match)
        }
    }
    return sum
}

func processWithConditionals(lines []string) int {
    sum := 0
    enabled := true
    re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
    for _, line := range lines {
        for _, match := range re.FindAllString(line, -1) {
            if match[2] == '(' {
                enabled = true
            } else if match[2] == 'n' {
                enabled = false
            } else if enabled && match[2] == 'l' {
                sum += mul(match)
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
