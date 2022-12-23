package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	inspected int
	items     []int
	operation func(int) int
	test      func(int) bool
	ifTrue    int
	ifFalse   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	monkeys := parseInput(scanner)

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			//fmt.Printf("Monkey %d\n", j)
			for _, old := range monkey.items {
				//fmt.Printf("  Monkey inspects an item with a worry level of %d\n", old)
				monkey.inspected++
				new := monkey.operation(old) / 3
				//fmt.Printf("    New item worry level %d\n", new)
				var newMon int
				if monkey.test(new) {
					//fmt.Printf("    Passes test\n")
					newMon = monkey.ifTrue
				} else {
					//fmt.Printf("    Fails test\n")
					newMon = monkey.ifFalse
				}
				monkeys[newMon].items = append(monkeys[newMon].items, new)
				//fmt.Printf("    Thrown to monkey %d\n", newMon)
			}
			monkey.items = []int{}
		}
		//printWorryLevels(monkeys)
	}

	inspections := []int{}

	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspected)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Printf("Result: %d\n", inspections[0]*inspections[1])

}

func printWorryLevels(monkeys []*monkey) {
	for i, monkey := range monkeys {
		fmt.Printf("Monkey %d: %v\n", i, monkey.items)
	}
}

func parseInput(scanner *bufio.Scanner) []*monkey {
	monkeys := []*monkey{}
	for scanner.Scan() {
		scanner.Scan()
		mon := monkey{}
		_, after, _ := strings.Cut(scanner.Text(), ": ")
		mon.items = map2(strings.Split(after, ", "), func(s string) (n int) {
			n, _ = strconv.Atoi(s)
			return
		})
		scanner.Scan()
		_, after, _ = strings.Cut(scanner.Text(), "new = ")
		opStrings := strings.Split(after, " ")
		var self bool
		val := 0
		if opStrings[0] == "old" && opStrings[2] == "old" {
			self = true
		} else {
			self = false
			val, _ = strconv.Atoi(opStrings[2])
		}
		mon.operation = createOpFunc(opStrings[1], val, self)
		scanner.Scan()
		_, after, _ = strings.Cut(scanner.Text(), "by ")
		val, _ = strconv.Atoi(after)
		mon.test = createTestFunc(val)
		scanner.Scan()
		_, after, _ = strings.Cut(scanner.Text(), "monkey ")
		val, _ = strconv.Atoi(after)
		mon.ifTrue = val
		scanner.Scan()
		_, after, _ = strings.Cut(scanner.Text(), "monkey ")
		val, _ = strconv.Atoi(after)
		mon.ifFalse = val
		monkeys = append(monkeys, &mon)
		scanner.Scan()
	}
	return monkeys
}

func createTestFunc(val int) func(int) bool {
	return func(n int) bool {
		return n%val == 0
	}
}

// Create new closure operation function
func createOpFunc(op string, val int, self bool) func(int) int {
	return func(old int) (new int) {
		if self {
			val = old
		}
		switch op {
		case "+":
			new = old + val
		case "-":
			new = old - val
		case "*":
			new = old * val
		case "/":
			new = old / val
		}
		return
	}
}

func map2(input []string, f func(string) int) []int {
	mapped := make([]int, len(input))
	for i, e := range input {
		mapped[i] = f(e)
	}
	return mapped
}
