package main

import (
	"fmt"
	"time"

	"github.com/sidmund/aoc-2024/lib"
)

func patrol(start lib.Point, obstruction lib.Point, lab [][]byte) map[lib.Point]int {
	delta := []lib.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}} // NESW
	pos, dir := start, 0
	visited := map[lib.Point]int{}

	for {
		if 1<<dir&visited[pos] != 0 {
			return nil // Guard's already been here, so he's stuck in a loop
		}

		visited[pos] |= 1 << dir // Mark this dir for this point as visited
		next := pos.Add(delta[dir])
		if next.Y < 0 || next.Y >= len(lab) || next.X < 0 || next.X >= len(lab[0]) {
			return visited // Guard leaves
		}

		if char := lab[next.Y][next.X]; char == '#' || next == obstruction {
			dir = (dir + 1) % 4 // Right turn
		} else {
			pos = next
		}
	}
}

func main() {
	lines := lib.ReadLines("day06/input")

	lab := make([][]byte, len(lines))
	var start lib.Point
	for y, line := range lines {
		lab[y] = []byte(line)
		for x, char := range lab[y] {
			if char == '^' {
				start = lib.Point{X: x, Y: y}
			}
		}
	}

	visited := patrol(start, lib.Point{X: -1, Y: -1}, lab)
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
