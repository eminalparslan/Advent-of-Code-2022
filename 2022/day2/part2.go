package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func contains(arr [3]string, i string) bool {
	for _, j := range arr {
		if i == j {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0

	rock := [3]string{"A Y", "B X", "C Z"}
	paper := [3]string{"A Z", "B Y", "C X"}

	for scanner.Scan() {
		round := scanner.Text()
		// Check win condition
		if round[2] == 'Y' {
			totalScore += 3
		} else if round[2] == 'Z' {
			totalScore += 6
		}
		// Check what I play
		if contains(rock, round) {
			totalScore += 1
		} else if contains(paper, round) {
			totalScore += 2
		} else {
			totalScore += 3
		}
	}

	fmt.Printf("Total score: %d\n", totalScore)
}
