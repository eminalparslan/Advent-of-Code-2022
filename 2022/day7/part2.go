package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	const (
		totalSpace  = 70_000_000
		neededSpace = 30_000_000
	)

	usedSpace := root.fileSize
	remainingSpace := totalSpace - usedSpace

	if remainingSpace >= neededSpace {
		fmt.Printf("Don't need to do anything")
		return
	}

	toFreeSpace := neededSpace - remainingSpace

	sizes := []int{}
	root.filter(func(n int) bool {
		return n >= toFreeSpace
	}, &sizes)

	min := sizes[0]
	for _, size := range sizes {
		if size < min {
			min = size
		}
	}

	fmt.Printf("Result: %d\n", min)
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
