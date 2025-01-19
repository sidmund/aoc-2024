package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sidmund/aoc-2024/lib"
)

func fragment(disk []int) int {
	defer lib.Measure(time.Now(), "Part 1")

	checksum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] < 0 {
			j := len(disk) - 1
			for j > i {
				if disk[j] > -1 {
					disk[i], disk[j] = disk[j], disk[i]
					break
				}
				j--
			}
			if j == i { // Only free space to the right
				break
			}
		}

		if disk[i] > -1 {
			checksum += i * disk[i]
		}
	}
	return checksum
}

func compact(disk []int) int {
	defer lib.Measure(time.Now(), "Part 2")

	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] < 0 {
			continue
		}

		size := 1
		for s := i - 1; s >= 0; s-- {
			if disk[s] == disk[i] {
				size++
			} else {
				break
			}
		}
		i -= size - 1

		for j := 0; j < i; j++ {
			if disk[j] < 0 {
				free := 1
				for f := j + 1; f < i; f++ {
					if disk[f] < 0 {
						free++
					} else {
						break
					}
				}
				j += free - 1

				if size <= free {
					for k := 0; k < size; k++ {
						disk[i+k], disk[j-(free-1)+k] = disk[j-(free-1)+k], disk[i+k]
					}
					break
				}
			}
		}
	}

	checksum := 0
	for i, id := range disk {
		if id > -1 {
			checksum += i * id
		}
	}
	return checksum
}

func main() {
	diskmap, _ := os.ReadFile("day09/input")

	disk1, disk2 := []int{}, []int{}
	for i := 0; i < len(diskmap)-1; i++ {
		blocks := int(diskmap[i]) - '0'
		if i%2 == 0 {
			id := i / 2
			for b := 0; b < blocks; b++ {
				disk1 = append(disk1, id)
				disk2 = append(disk2, id)
			}
		} else {
			for b := 0; b < blocks; b++ {
				disk1 = append(disk1, -1)
				disk2 = append(disk2, -1)
			}
		}
	}

	fmt.Println("Part 1", fragment(disk1))
	fmt.Println("Part 2", compact(disk2))
}
