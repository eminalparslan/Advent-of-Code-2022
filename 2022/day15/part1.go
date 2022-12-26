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
	file, err := os.Open("testInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		_, after, _ := strings.Cut(scanner.Text(), "Sensor at ")
		signal, beacon, _ := strings.Cut(after, ": closest beacon is at ")
		sX, sY := parseCoords(signal)
		bX, bY := parseCoords(beacon)
		dist := abs(sX-bX) + abs(sY-bY)
	}

	fmt.Printf("Hello world\n")
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
