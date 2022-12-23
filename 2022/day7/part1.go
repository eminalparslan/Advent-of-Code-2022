package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO: compare speed of pointer vs non-pointer dirs map

type dir struct {
	fileSize int
	dirs     map[string]*dir
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	root := dir{0, map[string]*dir{}}
	root.parse(scanner)
	root.accumulate()

	sizes := []int{}
	root.filter(func(n int) bool {
		return n <= 100_000
	}, &sizes)

	total := 0
	for _, size := range sizes {
		total += size
	}

	fmt.Printf("Result: %d\n", total)
}

func parseCommand(scanner *bufio.Scanner) []string {
	line := strings.TrimPrefix(scanner.Text(), "$ ")
	return strings.Split(line, " ")
}

func (curDir *dir) parse(scanner *bufio.Scanner) {
	scanner.Scan()
	command := parseCommand(scanner)
	if command[0] == "ls" {
		for scanner.Scan() && scanner.Text()[0] != '$' {
			files := strings.Split(scanner.Text(), " ")
			if files[0] == "dir" {
				curDir.dirs[files[1]] = &dir{0, map[string]*dir{}}
			} else {
				size, _ := strconv.Atoi(files[0])
				curDir.fileSize += size
			}
		}
		command = parseCommand(scanner)
	}
	for command[0] == "cd" {
		if command[1] == ".." {
			return
		} else if subDir, ok := curDir.dirs[command[1]]; ok {
			subDir.parse(scanner)
		}
		scanner.Scan()
		command = parseCommand(scanner)
	}
}

func (curDir *dir) accumulate() {
	for _, subDir := range curDir.dirs {
		subDir.accumulate()
		curDir.fileSize += subDir.fileSize
	}
}

func (curDir *dir) filter(f func(int) bool, sizes *[]int) {
	for _, subDir := range curDir.dirs {
		subDir.filter(f, sizes)
	}
	if f(curDir.fileSize) {
		*sizes = append(*sizes, curDir.fileSize)
	}
}
