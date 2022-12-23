package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cals := []int{}
	cur := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			cals = append(cals, cur)
			cur = 0
		} else {
			num, err := strconv.Atoi(scanner.Text())
			check(err)
			cur += num
		}
	}

	sort.Slice(cals, func(i, j int) bool {
		return cals[i] > cals[j]
	})

	//fmt.Printf("Cals: %v\n", cals)
	fmt.Printf("Total: %d\n", cals[0]+cals[1]+cals[2])

}
