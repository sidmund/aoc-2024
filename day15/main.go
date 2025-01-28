package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sidmund/aoc-2024/lib"
)

type box struct{ x, y, w int }

func boxAt(boxes map[box]bool, p lib.Point) (box, bool) {
	for b := range boxes {
		if p.Y == b.y && p.X >= b.x && p.X < b.x+b.w {
			return b, true
		}
	}
	return box{}, false
}

func push(boxes map[box]bool, walls map[lib.Point]bool, start, d lib.Point) bool {
	toPush, n := []box{}, start
	for {
		if b, ok := boxAt(boxes, n); ok {
			toPush = append(toPush, b)
			n = n.Add(lib.Point{X: b.w * d.X, Y: d.Y})
		} else {
			break
		}
	}
	if walls[n] {
		return false // Cannot push
	}
	if n == start {
		return true // Empty spot, no pushing required
	}
	for i := len(toPush) - 1; i >= 0; i-- {
		delete(boxes, toPush[i])
		boxes[box{toPush[i].x + d.X, toPush[i].y + d.Y, toPush[i].w}] = true
	}
	return true
}

func canPushVertical(boxes map[box]bool, walls map[lib.Point]bool, start lib.Point, dy int) bool {
	if walls[start] {
		return false
	}
	if b, ok := boxAt(boxes, start); ok {
		l, r := lib.Point{X: b.x, Y: b.y + dy}, lib.Point{X: b.x + 1, Y: b.y + dy}
		if walls[l] || walls[r] {
			return false
		}
		if canPushVertical(boxes, walls, l, dy) && canPushVertical(boxes, walls, r, dy) {
			return true
		}
		return false
	}
	return true // Empty spot
}

func pushVertical(boxes map[box]bool, walls map[lib.Point]bool, start lib.Point, dy int) {
	if b, ok := boxAt(boxes, start); ok {
		l, r := lib.Point{X: b.x, Y: b.y + dy}, lib.Point{X: b.x + 1, Y: b.y + dy}
		pushVertical(boxes, walls, l, dy)
		pushVertical(boxes, walls, r, dy)
		delete(boxes, b)
		boxes[box{b.x, b.y + dy, b.w}] = true
	}
}

func gps(warehouse, moves string, width int) int {
	robot, boxes, walls := lib.Point{}, map[box]bool{}, map[lib.Point]bool{}
	for y, s := range strings.Split(warehouse, "\n") {
		for x, r := range s {
			if r == '#' {
				for w := 0; w < width; w++ {
					walls[lib.Point{X: width*x + w, Y: y}] = true
				}
			} else if r == 'O' {
				boxes[box{width * x, y, width}] = true
			} else if r == '@' {
				robot = lib.Point{X: width * x, Y: y}
			}
		}
	}

	delta := map[rune]lib.Point{'^': {X: 0, Y: -1}, 'v': {X: 0, Y: 1}, '>': {X: 1, Y: 0}, '<': {X: -1, Y: 0}}
	for _, r := range moves {
		if d, ok := delta[r]; ok {
			n := robot.Add(d)
			if width > 1 && d.Y != 0 {
				if canPushVertical(boxes, walls, n, d.Y) {
					pushVertical(boxes, walls, n, d.Y)
					robot = n
				}
			} else if push(boxes, walls, n, d) {
				robot = n
			}
		}
	}

	gps := 0
	for b := range boxes {
		gps += 100*b.y + b.x
	}
	return gps
}

func main() {
	raw, _ := os.ReadFile("day15/input")
	input := strings.Split(string(raw), "\n\n")

	fmt.Println("Part 1:", gps(input[0], input[1], 1))
	fmt.Println("Part 2:", gps(input[0], input[1], 2))
}
