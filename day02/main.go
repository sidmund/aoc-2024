package main

import (
    "fmt"
    "strconv"
    "strings"

	lib "github.com/sidmund/aoc-2024/lib"
)

func parse(s string) int64 {
    result, _ := strconv.ParseInt(s, 10, 64)
    return result
}

func isReportSafe(levels []string, skipIndex int) bool {
    asc, desc := false, false

    start, end := 1, len(levels)
    if skipIndex == 0 {
        start = 2
    } else if skipIndex == len(levels) - 1 {
        end = len(levels) - 1
    }

    for i := start; i < end; i++ {
        var level int64
        if i - 1 == skipIndex {
            level = parse(levels[i-2])
        } else {
            level = parse(levels[i-1])
        }

        var next int64
        if i == skipIndex {
            next = parse(levels[i+1])
        } else {
            next = parse(levels[i])
        }

        diff := next - level

        if diff > 3 || diff < -3 {
            return false
        }

        if diff > 0 {
            asc = true
        } else if diff < 0 {
            desc = true
        } else {
            return false
        }

        if asc && desc { // direction changed
            return false
        }
    }

    return true
}

func main() {
    lines := lib.ReadLines("day02/input")

    safe, safeWithRemoval := 0, 0
    for _, line := range lines {
        levels := strings.Fields(line)

        if isReportSafe(levels, -1) {
            safe++
            continue
        }

        for i := 0; i < len(levels); i++ {
            if isReportSafe(levels, i) {
                safeWithRemoval++
                break
            }
        }
    }

    fmt.Println("Part 1:", safe)
    fmt.Println("Part 2:", safe + safeWithRemoval)
}
