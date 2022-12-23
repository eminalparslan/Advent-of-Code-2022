package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPriority := 0

	for scanner.Scan() {
		rucksack := scanner.Text()
		// Two sets to represent compartments
		compartment1 := map[byte]bool{}
		compartment2 := map[byte]bool{}
		for i := 0; i < len(rucksack)/2; i++ {
			compartment1[rucksack[i]] = true
			compartment2[rucksack[len(rucksack)-i-1]] = true
		}
		// Find the intersection of the sets
		for k := range compartment1 {
			if compartment2[k] {
				if k >= 65 && k <= 90 {
					totalPriority += int(k) - 38
				} else if k >= 97 && k <= 122 {
					totalPriority += int(k) - 96
				}
				break
			}
		}
	}

	fmt.Printf("Total priority: %d\n", totalPriority)

}
