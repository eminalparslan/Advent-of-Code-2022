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

	towers := [][]rune{}
	ok := scanner.Scan()

	for i := 1; i < len(scanner.Text()); i += 4 {
		towers = append(towers, []rune{})
	}

	lines := []string{}

	for ok && scanner.Text()[0] == '[' {
		lines = append(lines, scanner.Text())
		ok = scanner.Scan()
	}

	for i := len(lines) - 1; i >= 0; i-- {
		for j, r := range lines[i] {
			if j%4 == 1 && r != ' ' {
				towers[j/4] = append(towers[j/4], r)
			}
		}
	}

	scanner.Scan()

	for scanner.Scan() {
		toks := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(toks[1])
		from, _ := strconv.Atoi(toks[3])
		to, _ := strconv.Atoi(toks[5])
		from--
		to--
		towers[to] = append(towers[to], towers[from][len(towers[from])-num:]...)
		towers[from] = towers[from][:len(towers[from])-num]
	}

	fmt.Print("Result: ")
	for i := 0; i < len(towers); i++ {
		var pop rune
		pop, towers[i] = towers[i][len(towers[i])-1], towers[i][:len(towers[i])-1]
		fmt.Printf("%c", pop)
	}
	fmt.Print("\n")

}
