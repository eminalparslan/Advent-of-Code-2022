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

	win := [3]string{"A Y", "B Z", "C X"}
	tie := [3]string{"A X", "B Y", "C Z"}

	for scanner.Scan() {
		round := scanner.Text()
		// Points for decision
		switch round[2] {
		case 'X':
			totalScore += 1
		case 'Y':
			totalScore += 2
		case 'Z':
			totalScore += 3
		}
		// Check win conditions
		if contains(win, round) {
			totalScore += 6
		} else if contains(tie, round) {
			totalScore += 3
		}
	}

	fmt.Printf("Total score: %d\n", totalScore)
}
