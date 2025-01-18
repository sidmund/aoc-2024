package main

import (
	"fmt"

	"github.com/sidmund/aoc-2024/lib"
)

type point struct {
	x int
	y int
}

func main() {
	lines := lib.ReadLines("day08/input")

	freq := map[rune][]point{}
	for y, line := range lines {
		for x, r := range line {
			if r != '.' {
				freq[r] = append(freq[r], point{x, y})
			}
		}
	}

	resonant, harmonics := map[point]struct{}{}, map[point]struct{}{}
	for _, antennas := range freq {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}

				dx, dy := a2.x-a1.x, a2.y-a1.y
				for i := 0; ; i++ { // 0, because the antenna itself counts in the harmonics as well
					antinode := point{a2.x + (i * dx), a2.y + (i * dy)}
					if antinode.x < 0 || antinode.x >= len(lines) || antinode.y < 0 || antinode.y >= len(lines) {
						break
					}

					harmonics[antinode] = struct{}{}
					if i == 1 {
						resonant[antinode] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Println("Part 1:", len(resonant))
	fmt.Println("Part 2:", len(harmonics))
}
