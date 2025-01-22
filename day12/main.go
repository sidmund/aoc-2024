package main

import (
	"fmt"

	"github.com/sidmund/aoc-2024/lib"
)

type plot struct{ x, y int }

func bfs(garden map[plot]rune) (int, int) {
	visited := map[plot]bool{}
	price, discounted := 0, 0
	for p := range garden {
		if visited[p] {
			continue
		}
		visited[p] = true
		queue := []plot{p}
		area, perimeter, sides := 1, 0, 0 // Nr of sides is akin to nr of corners
		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			for _, d := range []plot{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				n := plot{c.x + d.x, c.y + d.y}
				if garden[c] != garden[n] {
					perimeter++
					outer := plot{c.x - d.y, c.y + d.x} // Check diagonally
					inner := plot{outer.x + d.x, outer.y + d.y}
					if garden[c] != garden[outer] || garden[c] == garden[inner] {
						sides++
					}
				} else if !visited[n] {
					queue = append(queue, n)
					visited[n] = true
					area++
				}
			}
		}
		price += area * perimeter
		discounted += area * sides
	}
	return price, discounted
}

func main() {
	lines := lib.ReadLines("day12/input")
	garden := map[plot]rune{}
	for y, line := range lines {
		for x, r := range line {
			garden[plot{x, y}] = r
		}
	}

	price, discounted := bfs(garden)
	fmt.Println("Part 1:", price)
	fmt.Println("Part 2:", discounted)
}
