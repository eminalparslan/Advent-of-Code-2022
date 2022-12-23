package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const knots = 10

	rope := [knots]Pos{}
	tailVisited := map[Pos]bool{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(line[1])
		for i := 0; i < num; i++ {
			switch line[0] {
			case "U":
				rope[0].y++
			case "D":
				rope[0].y--
			case "L":
				rope[0].x--
			case "R":
				rope[0].x++
			}
			for j := 1; j < knots; j++ {
				if rope[j-1].y > rope[j].y+1 {
					rope[j].y++
					if rope[j-1].x > rope[j].x {
						rope[j].x++
					} else if rope[j-1].x < rope[j].x {
						rope[j].x--
					}
				} else if rope[j-1].y < rope[j].y-1 {
					rope[j].y--
					if rope[j-1].x > rope[j].x {
						rope[j].x++
					} else if rope[j-1].x < rope[j].x {
						rope[j].x--
					}
				} else if rope[j-1].x < rope[j].x-1 {
					rope[j].x--
					if rope[j-1].y > rope[j].y {
						rope[j].y++
					} else if rope[j-1].y < rope[j].y {
						rope[j].y--
					}
				} else if rope[j-1].x > rope[j].x+1 {
					rope[j].x++
					if rope[j-1].y > rope[j].y {
						rope[j].y++
					} else if rope[j-1].y < rope[j].y {
						rope[j].y--
					}
				}
			}
			tailVisited[rope[knots-1]] = true
		}
	}

	fmt.Printf("Result: %d\n", len(tailVisited))
}
