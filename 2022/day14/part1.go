package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	caveHeight = 170
	caveWidth  = 100
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cave := [caveHeight][caveWidth]byte{}

	for y := range cave {
		for x := range cave[y] {
			cave[y][x] = '.'
		}
	}

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")
		prevX, prevY := 0, 0
		for i, coord := range coords {
			strX, strY, _ := strings.Cut(coord, ",")
			x, _ := strconv.Atoi(strX)
			y, _ := strconv.Atoi(strY)
			x -= 450
			if i > 0 {
				if x == prevX {
					minY, maxY := y, prevY
					if prevY < minY {
						minY, maxY = prevY, y
					}
					for ; minY <= maxY; minY++ {
						cave[minY][x] = '#'
					}
				} else if y == prevY {
					minX, maxX := x, prevX
					if prevX < minX {
						minX, maxX = prevX, x
					}
					for ; minX <= maxX; minX++ {
						cave[y][minX] = '#'
					}
				}
			}
			prevX, prevY = x, y
		}
	}

	steps := 0
	for {
		if sandFall(&cave) {
			break
		}
		steps++
	}

	//for i := range cave {
	//	for j := range cave[i] {
	//		fmt.Printf("%c ", cave[i][j])
	//	}
	//	fmt.Print("\n")
	//}

	fmt.Printf("Steps: %d\n", steps)
}

func sandFall(cave *[caveHeight][caveWidth]byte) bool {
	sandX, sandY := 50, 0
	for {
		if cave[sandY+1][sandX] == '.' {
			sandY++
		} else if cave[sandY+1][sandX-1] == '.' {
			sandY++
			sandX--
		} else if cave[sandY+1][sandX+1] == '.' {
			sandY++
			sandX++
		} else {
			break
		}
		if sandY == len(cave)-1 {
			return true
		}
	}
	cave[sandY][sandX] = 'o'
	return false
}
