package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	buffer, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := 0

	for i := 3; i < len(buffer); i++ {
		chars := map[byte]bool{}
		chars[buffer[i]] = true
		chars[buffer[i-1]] = true
		chars[buffer[i-2]] = true
		chars[buffer[i-3]] = true
		if len(chars) == 4 {
			res = i + 1
			break
		}
	}

	fmt.Printf("Result: %d\n", res)
}
