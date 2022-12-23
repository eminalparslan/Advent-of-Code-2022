package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		sections1 := strings.Split(pairs[0], "-")
		sections2 := strings.Split(pairs[1], "-")
		s11, _ := strconv.Atoi(sections1[0])
		s12, _ := strconv.Atoi(sections1[1])
		s21, _ := strconv.Atoi(sections2[0])
		s22, _ := strconv.Atoi(sections2[1])
		fmt.Printf("%d <= %d && %d >= %d ||\n", s11, s21, s12, s21)
		fmt.Printf("%d <= %d && %d >= %d\n\n", s11, s22, s12, s22)
		// https://nedbatchelder.com/blog/201310/range_overlap_in_two_compares.html
		if s12 >= s21 && s22 >= s11 {
			count++
		}
	}

	fmt.Printf("Count: %d\n", count)
}
