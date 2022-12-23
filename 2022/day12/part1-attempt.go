package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Failed attempt: tried using recursive backtracking, wayyyy too slow

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	heightmap := [][]byte{}
	visited := [][]byte{}

	var startRow int
	var startCol int

	row := 0
	for scanner.Scan() {
		heightmap = append(heightmap, []byte(scanner.Text()))
		visited = append(visited, make([]byte, len(scanner.Text())))
		index := strings.Index(scanner.Text(), "S")
		if index != -1 {
			startRow = row
			startCol = index
		}
		row++
	}

	result, found := fewestSteps(heightmap, visited, startRow, startCol, 0, 'a')
	if !found {
		log.Fatal("No solution\n")
	}

	fmt.Printf("Result: %d\n", result)

}

func fewestSteps(heightmap, visited [][]byte, row, col, steps int, height byte) (int, bool) {
	if row < 0 || col < 0 || row >= len(heightmap) || col >= len(heightmap[0]) {
		return -1, false
	} else if visited[row][col] == 1 {
		//fmt.Printf("Visited\n")
		return -1, false
	}
	newHeight := heightmap[row][col]
	if newHeight == 'S' {
		newHeight = 'a'
	} else if newHeight == 'E' && (height == 'y' || height == 'z') {
		return steps, true
	} else if newHeight > height+1 {
		//fmt.Printf("Height issue: %d > %d\n", heightmap[row][col], height+1)
		return -1, false
	}
	//fmt.Printf("Row: %d, col: %d\n", row, col)
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	visited[row][col] = 1
	min := -1
	for i := 0; i < 4; i++ {
		dr, dc := dirs[i][0], dirs[i][1]
		newSteps, found := fewestSteps(heightmap, visited, row+dr, col+dc, steps+1, newHeight)
		if found && (min == -1 || newSteps < min) {
			min = newSteps
		}
	}
	visited[row][col] = 0
	if min == -1 {
		return -1, false
	} else {
		return min, true
	}
}
