package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	max := 0
	cur := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			if cur > max {
				max = cur
			}
			cur = 0
		} else {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			cur += num
		}
	}

	fmt.Printf("Max: %d\n", max)

}
