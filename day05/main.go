package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func middlePage(update []string, rules map[string][]string, nFixes int) (int, int) {
	for p, page := range update {
		for i := p + 1; i < len(update); i++ {
			if !slices.Contains(rules[page], update[i]) {
				update[p], update[i] = update[i], update[p]
				return middlePage(update, rules, nFixes+1)
			}
		}
	}
	middle, _ := strconv.Atoi(update[len(update)/2])
	return middle, nFixes
}

func main() {
	raw, _ := os.ReadFile("day05/input")
	split := strings.Split(string(raw), "\n\n")

	rules := make(map[string][]string)
	for _, rule := range strings.Fields(split[0]) {
		var from, to string
		fmt.Sscanf(rule, "%2s|%2s", &from, &to)
		rules[from] = append(rules[from], to)
	}

	sumCorrect, sumFixed := 0, 0
	for _, update := range strings.Fields(split[1]) {
		if middle, nFixes := middlePage(strings.Split(update, ","), rules, 0); nFixes == 0 {
			sumCorrect += middle
		} else {
			sumFixed += middle
		}
	}

	fmt.Println("Part 1:", sumCorrect)
	fmt.Println("Part 2:", sumFixed)
}
