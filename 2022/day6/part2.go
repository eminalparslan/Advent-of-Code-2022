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
	distinct := 14

	for i := distinct - 1; i < len(buffer); i++ {
		chars := map[byte]bool{}
		for j := 0; j < distinct; j++ {
			chars[buffer[i-j]] = true
		}
		if len(chars) == distinct {
			res = i + 1
			break
		}
	}

	fmt.Printf("Result: %d\n", res)
}
