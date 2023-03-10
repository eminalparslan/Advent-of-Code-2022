package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Vertex struct {
	height   byte
	distance int
	visited  bool
	edges    []*Vertex
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

	lowest := [][2]int{}
	endRow, endCol := 0, 0

	graph := make([][]Vertex, len(heightmap))
	for r := range heightmap {
		graph[r] = make([]Vertex, len(heightmap[0]))
		for c, h := range heightmap[r] {
			if h == 'S' {
				graph[r][c].height = 'a'
				lowest = append(lowest, [2]int{r, c})
				graph[r][c].distance = maxIntVal
			} else if h == 'E' {
				graph[r][c].height = 'z'
				endRow, endCol = r, c
				graph[r][c].distance = 0
			} else {
				graph[r][c].height = h
				graph[r][c].distance = maxIntVal
				if h == 'a' {
					lowest = append(lowest, [2]int{r, c})
				}
			}
		}
	}

	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for r := range graph {
		for c := range graph[r] {
			for _, dir := range dirs {
				dr, dc := dir[0], dir[1]
				if r+dr >= 0 && r+dr < len(graph) && c+dc >= 0 && c+dc < len(graph[0]) {
					if graph[r][c].height-1 <= graph[r+dr][c+dc].height {
						graph[r][c].edges = append(graph[r][c].edges, &graph[r+dr][c+dc])
					}
				}
			}
		}
	}

	pq := []*Vertex{}
	pq = append(pq, &graph[endRow][endCol])

	var v *Vertex
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

	min := 0
	for i, low := range lowest {
		r, c := low[0], low[1]
		dist := graph[r][c].distance
		if i == 0 || dist < min {
			min = dist
		}
	}

	fmt.Printf("Min distance: %v\n", min)
}
