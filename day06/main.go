package main

import (
	"fmt"
	"time"

	"github.com/sidmund/aoc-2024/lib"
)

type point struct {
	x int
	y int
}

func patrol(start point, obstruction point, lab [][]byte) map[point]int {
	delta := []point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // NESW
	pos, dir := start, 0
	visited := map[point]int{}

	for {
		if 1<<dir&visited[pos] != 0 {
			return nil // Guard's already been here, so he's stuck in a loop
		}

		visited[pos] |= 1 << dir // Mark this dir for this point as visited
		next := point{pos.x + delta[dir].x, pos.y + delta[dir].y}
		if next.y < 0 || next.y >= len(lab) || next.x < 0 || next.x >= len(lab[0]) {
			return visited // Guard leaves
		}

		if char := lab[next.y][next.x]; char == '#' || next == obstruction {
			dir = (dir + 1) % 4 // Right turn
		} else {
			pos = next
		}
	}
}

func main() {
	lines := lib.ReadLines("day06/input")

	lab := make([][]byte, len(lines))
	var start point
	for y, line := range lines {
		lab[y] = []byte(line)
		for x, char := range lab[y] {
			if char == '^' {
				start = point{x, y}
			}
		}
	}

	visited := patrol(start, point{-1, -1}, lab)
	fmt.Println("Part 1:", len(visited))

	defer lib.Measure(time.Now(), "Part 2")
	stuck := 0
	for obstruction := range visited {
		if patrol(start, obstruction, lab) == nil {
			stuck++
		}
	}
	fmt.Println("Part 2:", stuck)
}
