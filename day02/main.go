package main

import (
    "fmt"
    "strconv"
    "strings"

	lib "github.com/sidmund/aoc-2024/lib"
)

func isReportSafe(levels []string) bool {
    asc, desc := false, false

    for i := 1; i < len(levels); i++ {
        level, _ := strconv.ParseInt(levels[i-1], 10, 64)
        next, _ := strconv.ParseInt(levels[i], 10, 64)
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

func isReportSafeWithRemoval(levels []string) bool {
    for i := 0; i < len(levels); i++ {
        dampened := make([]string, len(levels))
        copy(dampened, levels)
        if isReportSafe(append(dampened[:i], dampened[i+1:]...)) {
            return true
        }
    }
    return false
}

func main() {
    lines := lib.ReadLines("day02/input")

    safe, safeWithRemoval := 0, 0
    for _, line := range lines {
        levels := strings.Fields(line)
        if isReportSafe(levels) {
            safe++
        } else if isReportSafeWithRemoval(levels) {
            safeWithRemoval++
        }
    }

    fmt.Println("Part 1:", safe)
    fmt.Println("Part 2:", safe + safeWithRemoval)
}
