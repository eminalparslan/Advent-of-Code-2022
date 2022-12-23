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

	X := 1
	cycle := 1

	for scanner.Scan() {
		op := strings.Split(scanner.Text(), " ")
		switch op[0] {
		case "noop":
			incCycle(&cycle, X)
		case "addx":
			incCycle(&cycle, X)
			incCycle(&cycle, X)
			num, _ := strconv.Atoi(op[1])
			X += num
		}
	}

}

func incCycle(cycle *int, X int) {
	pos := (*cycle - 1) % 40
	if X == pos || X+1 == pos || X-1 == pos {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if *cycle%40 == 0 {
		fmt.Print("\n")
	}
	*cycle++
}
