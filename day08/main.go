package main

import (
	"fmt"

	"github.com/sidmund/aoc-2024/lib"
)

func main() {
	lines := lib.ReadLines("day08/input")

	freq := map[rune][]lib.Point{}
	for y, line := range lines {
		for x, r := range line {
			if r != '.' {
				freq[r] = append(freq[r], lib.Point{X: x, Y: y})
			}
		}
	}

	resonant, harmonics := map[lib.Point]struct{}{}, map[lib.Point]struct{}{}
	for _, antennas := range freq {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}

				d := a2.Sub(a1)
				for i := 0; ; i++ { // 0, because the antenna itself counts in the harmonics as well
					antinode := a2.Add(d.Scale(i))
					if antinode.X < 0 || antinode.X >= len(lines) || antinode.Y < 0 || antinode.Y >= len(lines) {
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
