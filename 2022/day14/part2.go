package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cave := map[point]bool{}

	highestY := 0
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")
		prevX, prevY := 0, 0
		for i, coord := range coords {
			strX, strY, _ := strings.Cut(coord, ",")
			x, _ := strconv.Atoi(strX)
			y, _ := strconv.Atoi(strY)
			if y > highestY {
				highestY = y
			}
			if i > 0 {
				if x == prevX {
					minY, maxY := y, prevY
					if prevY < minY {
						minY, maxY = prevY, y
					}
					for ; minY <= maxY; minY++ {
						cave[point{x: x, y: minY}] = true
					}
				} else if y == prevY {
					minX, maxX := x, prevX
					if prevX < minX {
						minX, maxX = prevX, x
					}
					for ; minX <= maxX; minX++ {
						cave[point{x: minX, y: y}] = true
					}
				}
			}
			prevX, prevY = x, y
		}
	}

	steps := 1
	for {
		if sandFall(&cave, highestY) {
			break
		}
		steps++
	}

	fmt.Printf("Steps: %d\n", steps)
}

func sandFall(cave *map[point]bool, highestY int) bool {
	sandX, sandY := 500, 0
	for {
		if sandY+1 == highestY+2 {
			break
		}
		if _, ok := (*cave)[point{x: sandX, y: sandY + 1}]; !ok {
			sandY++
		} else if _, ok := (*cave)[point{x: sandX - 1, y: sandY + 1}]; !ok {
			sandY++
			sandX--
		} else if _, ok := (*cave)[point{x: sandX + 1, y: sandY + 1}]; !ok {
			sandY++
			sandX++
		} else {
			break
		}
	}
	(*cave)[point{x: sandX, y: sandY}] = true
	if sandX == 500 && sandY == 0 {
		return true
	}
	return false
}
