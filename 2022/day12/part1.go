package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type vertex struct {
	height   byte
	distance int
	visited  bool
	edges    []*vertex
}

const maxIntVal = int(^uint(0) >> 1)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	heightmap := [][]byte{}

	for scanner.Scan() {
		heightmap = append(heightmap, []byte(scanner.Text()))
	}

	startRow, startCol := 0, 0
	endRow, endCol := 0, 0

	graph := make([][]vertex, len(heightmap))
	for r := range heightmap {
		graph[r] = make([]vertex, len(heightmap[0]))
		for c, h := range heightmap[r] {
			if h == 'S' {
				graph[r][c].height = 'a'
				startRow, startCol = r, c
				graph[r][c].distance = 0
			} else if h == 'E' {
				graph[r][c].height = 'z'
				endRow, endCol = r, c
				graph[r][c].distance = maxIntVal
			} else {
				graph[r][c].height = h
				graph[r][c].distance = maxIntVal
			}
		}
	}

	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for r := range graph {
		for c := range graph[r] {
			for _, dir := range dirs {
				dr, dc := dir[0], dir[1]
				if r+dr >= 0 && r+dr < len(graph) && c+dc >= 0 && c+dc < len(graph[0]) {
					if graph[r][c].height >= graph[r+dr][c+dc].height-1 {
						graph[r][c].edges = append(graph[r][c].edges, &graph[r+dr][c+dc])
					}
				}
			}
		}
	}

	pq := []*vertex{}
	pq = append(pq, &graph[startRow][startCol])

	var v *vertex
	for len(pq) != 0 {
		v, pq = pq[0], pq[1:]
		for _, e := range v.edges {
			if !e.visited && e.distance > v.distance+1 {
				e.distance = v.distance + 1
				pq = append(pq, e)
			}
		}
		v.visited = true
	}

	fmt.Printf("Distance: %v\n", graph[endRow][endCol].distance)
}
