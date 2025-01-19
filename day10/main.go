package main

import (
	"fmt"

	"github.com/sidmund/aoc-2024/lib"
)

type point struct{ x, y int }

func dfs(heights [][]int, start point, reached map[point]struct{}) int {
	if heights[start.y][start.x] == 9 {
		if _, ok := reached[start]; ok {
			return 0
		}
		if reached != nil {
			reached[start] = struct{}{}
		}
		return 1 // New trail found
	}

	score := 0
	for _, d := range []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		next := point{start.x + d.x, start.y + d.y}
		if next.x < 0 || next.x >= len(heights[0]) || next.y < 0 || next.y >= len(heights) {
			continue
		}
		if heights[next.y][next.x] == heights[start.y][start.x]+1 {
			score += dfs(heights, next, reached)
		}
	}
	return score
}

func main() {
	lines := lib.ReadLines("day10/input")
	heights, heads := make([][]int, len(lines)), []point{}
	for y, line := range lines {
		heights[y] = make([]int, len(line))
		for x, r := range line {
			heights[y][x] = int(r - '0')
			if r == '0' {
				heads = append(heads, point{x, y})
			}
		}
	}

	scores, ratings := 0, 0
	for _, head := range heads {
		scores += dfs(heights, head, map[point]struct{}{})
		ratings += dfs(heights, head, nil)
	}
	fmt.Println("Part 1", scores)
	fmt.Println("Part 2", ratings)
}
