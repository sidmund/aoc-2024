package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct{ x, y int }

func (p point) add(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

type box struct{ x, y, w int }

func boxAt(boxes map[box]bool, p point) (box, bool) {
	for b := range boxes {
		if p.y == b.y && p.x >= b.x && p.x < b.x+b.w {
			return b, true
		}
	}
	return box{}, false
}

func push(boxes map[box]bool, walls map[point]bool, start, d point) bool {
	toPush, n := []box{}, start
	for {
		if b, ok := boxAt(boxes, n); ok {
			toPush = append(toPush, b)
			n = n.add(point{b.w * d.x, d.y})
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
		boxes[box{toPush[i].x + d.x, toPush[i].y + d.y, toPush[i].w}] = true
	}
	return true
}

func canPushVertical(boxes map[box]bool, walls map[point]bool, start point, dy int) bool {
	if walls[start] {
		return false
	}
	if b, ok := boxAt(boxes, start); ok {
		l, r := point{b.x, b.y + dy}, point{b.x + 1, b.y + dy}
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

func pushVertical(boxes map[box]bool, walls map[point]bool, start point, dy int) {
	if b, ok := boxAt(boxes, start); ok {
		l, r := point{b.x, b.y + dy}, point{b.x + 1, b.y + dy}
		pushVertical(boxes, walls, l, dy)
		pushVertical(boxes, walls, r, dy)
		delete(boxes, b)
		boxes[box{b.x, b.y + dy, b.w}] = true
	}
}

func gps(warehouse, moves string, width int) int {
	robot, boxes, walls := point{}, map[box]bool{}, map[point]bool{}
	for y, s := range strings.Split(warehouse, "\n") {
		for x, r := range s {
			if r == '#' {
				for w := 0; w < width; w++ {
					walls[point{width*x + w, y}] = true
				}
			} else if r == 'O' {
				boxes[box{width * x, y, width}] = true
			} else if r == '@' {
				robot = point{width * x, y}
			}
		}
	}

	delta := map[rune]point{'^': {0, -1}, 'v': {0, 1}, '>': {1, 0}, '<': {-1, 0}}
	for _, r := range moves {
		if d, ok := delta[r]; ok {
			n := robot.add(d)
			if width > 1 && d.y != 0 {
				if canPushVertical(boxes, walls, n, d.y) {
					pushVertical(boxes, walls, n, d.y)
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
