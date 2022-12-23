package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("testInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		left := scanner.Text()
		scanner.Scan()
		right := scanner.Scan()

		scanner.Scan()
	}

	fmt.Printf("Hello world\n")
}
