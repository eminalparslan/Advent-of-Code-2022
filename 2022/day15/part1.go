package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// set for x's at y = 2000000 that can't contain a beacon
	excluded := map[int]bool{}
	// set of sensors and beacons
	sensorsAndBeacons := map[int]bool{}

	const targetY = 2_000_000

	for scanner.Scan() {
		_, after, _ := strings.Cut(scanner.Text(), "Sensor at ")
		sensor, beacon, _ := strings.Cut(after, ": closest beacon is at ")
		sX, sY := parseCoords(sensor)
		bX, bY := parseCoords(beacon)
		if sY == targetY {
			sensorsAndBeacons[sX] = true
		}
		if bY == targetY {
			sensorsAndBeacons[bX] = true
		}
		dist := abs(sX-bX) + abs(sY-bY)
		diffY := abs(sY - targetY)
		if diffY <= dist {
			n := dist - diffY
			for i := 0; i <= n; i++ {
				excluded[sX+i] = true
				excluded[sX-i] = true
			}
		}
	}

	res := len(excluded)

	for k := range excluded {
		if _, ok := sensorsAndBeacons[k]; ok {
			res--
		}
	}

	fmt.Printf("Result: %d\n", res)
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
