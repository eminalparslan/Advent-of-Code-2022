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
	total := 0

	for scanner.Scan() {
		op := strings.Split(scanner.Text(), " ")
		switch op[0] {
		case "noop":
			incCycle(&total, &cycle, X)
		case "addx":
			incCycle(&total, &cycle, X)
			num, _ := strconv.Atoi(op[1])
			X += num
			incCycle(&total, &cycle, X)
		}
	}

	fmt.Printf("Total: %d\n", total)

}

func incCycle(total, cycle *int, X int) {
	*cycle++
	if *cycle >= 20 && (*cycle-20)%40 == 0 {
		*total += *cycle * X
	}
}
