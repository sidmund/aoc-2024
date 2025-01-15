package main

import (
    "fmt"

    "github.com/sidmund/aoc-2024/lib"
)

func char(lines []string, i int, j int) byte {
    if j < 0 || j > len(lines)-1 || i < 0 || i > len(lines[j])-1 {
        return '.'
    }
    return lines[j][i]
}

func findMAS(lines []string, x int, y int) int {
    count := 0
    for dy := -1; dy < 2; dy++ {
        for dx := -1; dx < 2; dx++ {
            if dy == 0 && dx == 0 {
                continue
            }

            if char(lines, x+dx, y+dy) == 'M' && char(lines, x+dx*2, y+dy*2) == 'A' && char(lines, x+dx*3, y+dy*3) == 'S' {
                count++
            }
        }
    }
    return count
}

func isX(lines []string, x int, y int) bool {
    return (
        (char(lines, x-1, y-1) == 'M' && char(lines, x+1, y+1) == 'S') ||
        (char(lines, x-1, y-1) == 'S' && char(lines, x+1, y+1) == 'M')) &&
    (
        (char(lines, x+1, y-1) == 'M' && char(lines, x-1, y+1) == 'S') ||
        (char(lines, x+1, y-1) == 'S' && char(lines, x-1, y+1) == 'M'))
}

func main() {
    lines := lib.ReadLines("day04/input")

    count1, count2 := 0, 0
    for y, line := range lines {
        for x, char := range line {
            if char == 'X' {
                count1 += findMAS(lines, x, y)
            } else if char == 'A' && isX(lines, x, y) {
                count2++
            }
        }
    }

    fmt.Println("Part 1:", count1)
    fmt.Println("Part 2:", count2)
}
