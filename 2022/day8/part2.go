package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const size = 99

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	trees := make([][]byte, size)

	for i := 0; i < size; i++ {
		trees[i] = make([]byte, size+1)
		_, err := file.Read(trees[i])
		if err == io.EOF {
			log.Fatal("Reached end of file\n")
		}
		// Get rid of null byte
		trees[i] = trees[i][:size]
	}

	// Outer ring will always have 0 scenic score
	max := 0
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			score := 1
			k := 1
			for j+k < size && trees[i][j] > trees[i][j+k] {
				k++
			}
			if j+k == size {
				k--
			}
			score *= k
			k = 1
			for j-k >= 0 && trees[i][j] > trees[i][j-k] {
				k++
			}
			if j-k == -1 {
				k--
			}
			score *= k
			k = 1
			for i+k < size && trees[i][j] > trees[i+k][j] {
				k++
			}
			if i+k == size {
				k--
			}
			score *= k
			k = 1
			for i-k >= 0 && trees[i][j] > trees[i-k][j] {
				k++
			}
			if i-k == -1 {
				k--
			}
			score *= k
			if score > max {
				max = score
			}
		}
	}

	fmt.Printf("Highest: %d\n", max)
}
