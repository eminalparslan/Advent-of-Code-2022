package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

///////// Vertex struct definition ///////////////////

type vertex struct {
	name      string          // name of the valve
	rate      int             // flow rate of valve
	visited   bool            // whether the valve has been visited (used for dijkstra algorithm)
	distance  int             // distance used for dijkstra algorithm
	index     int             // needed for priority queue implementation
	edges     map[*vertex]int // distances to vertices connected by edges
	distances map[*vertex]int // distances to every other vertex
}

///////// Priority queue implementation ///////////////////

type priorityQueue []*vertex

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	v := x.(*vertex)
	v.index = n
	*pq = append(*pq, v)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	v := old[n-1]
	old[n-1] = nil
	v.index = -1
	*pq = old[0 : n-1]
	return v
}

const maxIntVal = int(^uint(0) >> 1)

///////// Main ///////////////////

func main() {
	file, err := os.Open("testInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// stores mapping from vertex name to vertex struct
	// nice way to access all vertices
	vertices := map[string]*vertex{}
	var start *vertex

	for scanner.Scan() {
		// parse file contents
		_, after, _ := strings.Cut(scanner.Text(), "Valve ")
		name, after, _ := strings.Cut(after, " has flow rate=")
		rateStr, conStr, found := strings.Cut(after, "; tunnels lead to valves ")
		// deal with single connection case
		if !found {
			rateStr, conStr, _ = strings.Cut(after, "; tunnel leads to valve ")
		}
		rate, _ := strconv.Atoi(rateStr)
		var connections []string
		// deal with single connection case here as well
		if len(conStr) == 2 {
			connections = append(connections, conStr)
		} else {
			connections = strings.Split(conStr, ", ")
		}
		// initialize vertex struct for each valve
		curVertex := &vertex{
			name:      name,
			rate:      rate,
			visited:   false,
			distance:  maxIntVal,
			index:     0,
			edges:     map[*vertex]int{},
			distances: map[*vertex]int{},
		}
		// create two way connection between this vertex and all connected vertices
		for _, con := range connections {
			if conVertex, ok := vertices[con]; ok {
				curVertex.edges[conVertex] = 1
				conVertex.edges[curVertex] = 1
			}
		}
		vertices[name] = curVertex
		if name == "AA" {
			start = curVertex
		}
	}

	// precomputation
	condenseGraph(&vertices)
	calculateDistances(&vertices)

	// to reset visited field
	resetDistances(&vertices)

	res1 := solve1(start, 30)
	res2 := solve2(start, start, 26, 26)
	max := res1
	if res2 > max {
		max = res2
	}
	fmt.Printf("Result: %d\n", max)

	// for _, v := range vertices {
	// 	fmt.Printf("%s: ", v.name)
	// 	for e, d := range v.edges {
	// 		fmt.Printf("%s:%d, ", e.name, d)
	// 	}
	// 	fmt.Print("distances: ")
	// 	for e, d := range v.distances {
	// 		fmt.Printf("%s:%d, ", e.name, d)
	// 	}
	// 	fmt.Print("\n")
	// }

}

// condenses the graph by removing 0-flow rate vertices and replacing them with edge weights
func condenseGraph(vertices *map[string]*vertex) {
	for k, v := range *vertices {
		// find all 0-flow rate vertices (except start)
		if v.rate == 0 && v.name != "AA" {
			// go through all its edge vertices
			for e1, d1 := range v.edges {
				// remove their connection to this vertex
				delete(e1.edges, v)
				// connect all edge vertices with every other edge vertex, summing edge weights
				for e2, d2 := range v.edges {
					if e1 != e2 {
						e1.edges[e2] = d1 + d2
						e2.edges[e1] = d1 + d2
					}
				}
			}
			// remove this 0-flow rate vertex
			delete(*vertices, k)
		}
	}
}

// resets distances and visited markers for all vertices in between dijkstra runs
func resetDistances(vertices *map[string]*vertex) {
	for _, v := range *vertices {
		v.distance = maxIntVal
		v.visited = false
	}
}

func calculateDistances(vertices *map[string]*vertex) {
	// calculate dijkstra shortest map to all vertices from every vertex
	for _, startVertex := range *vertices {
		resetDistances(vertices)
		startVertex.distance = 0
		pq := priorityQueue{startVertex}
		heap.Init(&pq)
		for len(pq) != 0 {
			curVertex := heap.Pop(&pq).(*vertex)
			for e := range curVertex.edges {
				newDistance := curVertex.distance + curVertex.edges[e]
				if !e.visited && e.distance > newDistance {
					e.distance = newDistance
					heap.Push(&pq, e)
				}
			}
			curVertex.visited = true
		}
		// add all shortest distances to this vertex's distances field
		for _, curVertex := range *vertices {
			if curVertex != startVertex {
				startVertex.distances[curVertex] = curVertex.distance
			}
		}
	}
}

func solve1(curVertex *vertex, timeLeft int) int {
	// base case: return 0 when time is out or vertex was already visited
	if timeLeft <= 0 || curVertex.visited {
		return 0
	}

	// calculate pressure released from this valve
	pressure := curVertex.rate * timeLeft

	curVertex.visited = true
	maxPressure := 0
	// iterte through all other vertices, recursively finding maximum pressure released
	for nextVertex, distance := range curVertex.distances {
		nextPressure := solve1(nextVertex, timeLeft-distance-1)
		if nextPressure > maxPressure {
			maxPressure = nextPressure
		}
	}
	curVertex.visited = false

	pressure += maxPressure
	return pressure
}

// takes like ~5 mins to calculate lol
func solve2(vertex1, vertex2 *vertex, timeLeft1, timeLeft2 int) int {
	if timeLeft1 <= 0 || vertex1.visited || timeLeft2 <= 0 || vertex2.visited {
		return 0
	}

	pressure := vertex1.rate*timeLeft1 + vertex2.rate*timeLeft2

	vertex1.visited = true
	vertex2.visited = true
	maxPressure := 0
	for nextVertex1, distance1 := range vertex1.distances {
		for nextVertex2, distance2 := range vertex2.distances {
			if nextVertex1 != nextVertex2 {
				nextPressure := solve2(nextVertex1, nextVertex2, timeLeft1-distance1-1, timeLeft2-distance2-1)
				if nextPressure > maxPressure {
					maxPressure = nextPressure
				}
			}
		}
	}
	vertex1.visited = false
	vertex2.visited = false

	pressure += maxPressure
	return pressure
}
