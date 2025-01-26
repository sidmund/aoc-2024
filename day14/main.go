package main

import (
	"fmt"
	"math"

	"github.com/sidmund/aoc-2024/lib"
)

var width, height int = 101, 103
var xAxis, yAxis int = width / 2, height / 2

type robot struct{ x, y, dx, dy int }

func (r *robot) move() robot {
	r.x = (r.x + r.dx) % width
	r.y = (r.y + r.dy) % height
	if r.x < 0 {
		r.x = width + r.x
	}
	if r.y < 0 {
		r.y = height + r.y
	}
	return *r
}

func main() {
	lines := lib.ReadLines("day14/input")
	robots := make([]robot, len(lines))
	for i, line := range lines {
		var r robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.x, &r.y, &r.dx, &r.dy)
		robots[i] = r
	}

	seconds := 10000
	scores, sum := make([]int, seconds), 0
	for t := 0; t < seconds; t++ {
		q1, q2, q3, q4 := 0, 0, 0, 0
		for i, r := range robots {
			robots[i] = r.move()
			if r.x < xAxis && r.y < yAxis {
				q1++
			} else if r.x > xAxis && r.y < yAxis {
				q2++
			} else if r.x > xAxis && r.y > yAxis {
				q3++
			} else if r.x < xAxis && r.y > yAxis {
				q4++
			}
		}
		scores[t] = q1 * q2 * q3 * q4
		sum += scores[t]
		if t == 99 {
			fmt.Println("Part 1:", scores[t])
		}
	}

	avg, squaredSum := sum/seconds, 0
	for _, score := range scores {
		squaredSum += (score - avg) * (score - avg)
	}
	std := int(math.Sqrt(float64(squaredSum / seconds)))
	for t, score := range scores {
		if score < avg-8*std || score > avg+8*std {
			fmt.Printf("Part 2: %d (outlier score of %d)\n", t+1, score)
			break
		}
	}
}
