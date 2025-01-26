package main

import (
	"fmt"
	"os"
	"strings"
)

func cramer(a, b, c, d, x, y int) int {
	if A, B := (d*x-b*y)/(a*d-b*c), (a*y-c*x)/(a*d-b*c); a*A+b*B == x && c*A+d*B == y {
		return 3*A + B
	}
	return 0
}

func main() {
	raw, _ := os.ReadFile("day13/input")
	tokens1, tokens2 := 0, 0
	for _, eqn := range strings.Split(string(raw), "\n\n") {
		var a, c, b, d, x, y int
		fmt.Sscanf(eqn, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &a, &c, &b, &d, &x, &y)
		tokens1 += cramer(a, b, c, d, x, y)
		tokens2 += cramer(a, b, c, d, x+10000000000000, y+10000000000000)
	}
	fmt.Println("Part 1:", tokens1)
	fmt.Println("Part 2:", tokens2)
}
