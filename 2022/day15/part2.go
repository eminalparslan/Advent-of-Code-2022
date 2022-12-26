package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const limit = 4_000_000 + 1

type bound struct {
	// left and right x bounds (inclusive)
	lX, rX int
}

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

	// list of bounds for each row
	rows := [limit][]bound{}
	beacons := map[point]bool{}

	for scanner.Scan() {
		_, after, _ := strings.Cut(scanner.Text(), "Sensor at ")
		sensor, beacon, _ := strings.Cut(after, ": closest beacon is at ")
		sX, sY := parseCoords(sensor)
		bX, bY := parseCoords(beacon)
		beacons[point{bX, bY}] = true
		dist := abs(sX-bX) + abs(sY-bY)
		for dY := 0; dY <= dist; dY++ {
			if sY+dY < limit {
				rows[sY+dY] = append(rows[sY+dY], bound{lX: sX - dist + dY, rX: sX + dist - dY})
			}
			if sY-dY >= 0 {
				rows[sY-dY] = append(rows[sY-dY], bound{lX: sX - dist + dY, rX: sX + dist - dY})
			}
		}
	}

	resX, resY := solve(rows, beacons)

	res := resX*(limit-1) + resY

	fmt.Printf("Result: %d\n", res)

}

func solve(rows [limit][]bound, beacons map[point]bool) (int64, int64) {
	for y, bounds := range rows {
		sort.Slice(bounds, func(i, j int) bool {
			return bounds[i].lX < bounds[j].lX
		})
		for l := 0; l < bounds[0].lX; l++ {
			if _, ok := beacons[point{l, y}]; !ok {
				return int64(l), int64(y)
			}
		}
		curBound := bounds[0]
		for _, b := range bounds {
			if b.lX > curBound.rX {
				for l := curBound.rX + 1; l < b.lX; l++ {
					if _, ok := beacons[point{l, y}]; !ok {
						return int64(l), int64(y)
					}
				}
			} else if b.rX > curBound.rX {
				curBound.rX = b.rX
			}
		}
	}
	return -1, -1
}

func parseCoords(s string) (x, y int) {
	coords := strings.Split(s, ", ")
	x, _ = strconv.Atoi(coords[0][2:])
	y, _ = strconv.Atoi(coords[1][2:])
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
