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
	visible := make([][]byte, size)

	for i := 0; i < size; i++ {
		visible[i] = make([]byte, size)

		trees[i] = make([]byte, size+1)
		_, err := file.Read(trees[i])
		if err == io.EOF {
			log.Fatal("Reached end of file\n")
		}
		trees[i] = trees[i][:size]
	}

	for i := 0; i < size; i++ {
		visible[i][0] = 1
		visible[0][i] = 1
		visible[i][size-1] = 1
		visible[size-1][i] = 1
	}

	for i := 1; i < size-1; i++ {
		maxSoFar := trees[i][0]
		for j := 1; j < size-1; j++ {
			if trees[i][j] > maxSoFar {
				visible[i][j] = 1
				maxSoFar = trees[i][j]
			}
		}
		maxSoFar = trees[i][size-1]
		for j := size - 2; j > 0; j-- {
			if trees[i][j] > maxSoFar {
				visible[i][j] = 1
				maxSoFar = trees[i][j]
			}
		}
	}

	for i := 1; i < size-1; i++ {
		maxSoFar := trees[0][i]
		for j := 1; j < size-1; j++ {
			if trees[j][i] > maxSoFar {
				visible[j][i] = 1
				maxSoFar = trees[j][i]
			}
		}
		maxSoFar = trees[size-1][i]
		for j := size - 2; j > 0; j-- {
			if trees[j][i] > maxSoFar {
				visible[j][i] = 1
				maxSoFar = trees[j][i]
			}
		}
	}

	total := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			total += int(visible[i][j])
		}
	}

	fmt.Printf("Total: %d\n", total)
}
