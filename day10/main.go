package main

import (
	"fmt"

	"github.com/sidmund/aoc-2024/lib"
)

func dfs(heights [][]int, start lib.Point, reached map[lib.Point]struct{}) int {
	if heights[start.Y][start.X] == 9 {
		if _, ok := reached[start]; ok {
			return 0
		}
		if reached != nil {
			reached[start] = struct{}{}
		}
		return 1 // New trail found
	}

	score := 0
	for _, d := range []lib.Point{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}} {
		next := start.Add(d)
		if next.X < 0 || next.X >= len(heights[0]) || next.Y < 0 || next.Y >= len(heights) {
			continue
		}
		if heights[next.Y][next.X] == heights[start.Y][start.X]+1 {
			score += dfs(heights, next, reached)
		}
	}
	return score
}

func main() {
	lines := lib.ReadLines("day10/input")
	heights, heads := make([][]int, len(lines)), []lib.Point{}
	for y, line := range lines {
		heights[y] = make([]int, len(line))
		for x, r := range line {
			heights[y][x] = int(r - '0')
			if r == '0' {
				heads = append(heads, lib.Point{X: x, Y: y})
			}
		}
	}

	scores, ratings := 0, 0
	for _, head := range heads {
		scores += dfs(heights, head, map[lib.Point]struct{}{})
		ratings += dfs(heights, head, nil)
	}
	fmt.Println("Part 1:", scores)
	fmt.Println("Part 2:", ratings)
}
