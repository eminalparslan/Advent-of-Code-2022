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

	headPos := Pos{0, 0}
	tailPos := Pos{0, 0}
	tailVisited := map[Pos]bool{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(line[1])
		for i := 0; i < num; i++ {
			switch line[0] {
			case "U":
				headPos.y++
				if headPos.y > tailPos.y+1 {
					tailPos.y++
					tailPos.x = headPos.x
				}
			case "D":
				headPos.y--
				if headPos.y < tailPos.y-1 {
					tailPos.y--
					tailPos.x = headPos.x
				}
			case "L":
				headPos.x--
				if headPos.x < tailPos.x-1 {
					tailPos.x--
					tailPos.y = headPos.y
				}
			case "R":
				headPos.x++
				if headPos.x > tailPos.x+1 {
					tailPos.x++
					tailPos.y = headPos.y
				}
			}
			tailVisited[tailPos] = true
		}
	}

	fmt.Printf("Result: %d\n", len(tailVisited))
}
