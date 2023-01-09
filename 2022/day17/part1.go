package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// input, err := os.ReadFile("testInput.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	file, err := os.Open("rocks.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rocks := [][4][4]byte{}

	i := 0
	rock := 0
	for scanner.Scan() {
		if i == 0 {
			rocks = append(rocks, [4][4]byte{})
		}
		for j := 0; j < 4; j++ {
			rocks[rock][i][j] = scanner.Text()[j]
		}
		i++
		if i == 4 {
			i = 0
			rock++
			scanner.Scan()
		}
	}

	fmt.Printf("Hello world: %v\n", rocks)
}
