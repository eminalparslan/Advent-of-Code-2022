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
		rucksack1 := scanner.Text()
		scanner.Scan()
		rucksack2 := scanner.Text()
		scanner.Scan()
		rucksack3 := scanner.Text()
		elf1 := map[byte]bool{}
		elf2 := map[byte]bool{}
		elf3 := map[byte]bool{}
		for _, c := range []byte(rucksack1) {
			elf1[c] = true
		}
		for _, c := range []byte(rucksack2) {
			elf2[c] = true
		}
		for _, c := range []byte(rucksack3) {
			elf3[c] = true
		}
		// Find the intersection of the sets
		for k := range elf1 {
			if elf2[k] && elf3[k] {
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
