package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type list struct {
	val      int
	children []*list
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		advance = 0
		// skip commas
		if data[0] == ',' {
			data = data[1:]
			advance++
		}

		if data[0] == '[' || data[0] == ']' {
			advance++
			token = data[:1]
		} else if data[0] >= 48 && data[0] <= 57 {
			i := 0
			for i < len(data) && data[i] >= 48 && data[i] <= 57 {
				i++
			}
			advance += i
			token = data[:i]
		} else {
			return 0, nil, errors.New("Unparseable input")
		}
		return
	}

	pairs := []*list{
		{
			val: -2, children: []*list{
				{
					val: -2, children: []*list{
						{val: 2},
					},
				},
			},
		},
		{
			val: -2, children: []*list{
				{
					val: -2, children: []*list{
						{val: 6},
					},
				},
			},
		},
	}

	for scanner.Scan() {
		left := bufio.NewScanner(strings.NewReader(scanner.Text()))
		left.Split(split)
		left.Scan()
		pairs = append(pairs, parse(left))
		scanner.Scan()
		right := bufio.NewScanner(strings.NewReader(scanner.Text()))
		right.Split(split)
		right.Scan()
		pairs = append(pairs, parse(right))
		scanner.Scan()
	}

	sort.Slice(pairs, func(i, j int) bool {
		res, _ := compare(pairs[i], pairs[j])
		return res
	})

	res := 1
	for i, l := range pairs {
		if l.val == -2 {
			res *= i + 1
		}
	}

	fmt.Printf("Result: %d\n", res)
}

func parse(scanner *bufio.Scanner) *list {
	tok := scanner.Bytes()[0]
	if tok == '[' {
		l := list{val: -1}
		scanner.Scan()
		tok = scanner.Bytes()[0]
		for tok != ']' {
			l.children = append(l.children, parse(scanner))
			scanner.Scan()
			tok = scanner.Bytes()[0]
		}
		return &l
	} else {
		num, _ := strconv.Atoi(scanner.Text())
		return &list{val: num}
	}
}

func compare(left, right *list) (bool, bool) {
	if len(left.children) == 0 && len(right.children) == 0 {
		return left.val < right.val, left.val != right.val
	} else if len(left.children) == 0 {
		left.children = append(left.children, &list{val: left.val})
		return compare(left, right)
	} else if len(right.children) == 0 {
		right.children = append(right.children, &list{val: right.val})
		return compare(left, right)
	} else {
		minLen := len(right.children)
		if len(left.children) < len(right.children) {
			minLen = len(left.children)
		}
		for i := 0; i < minLen; i++ {
			if res, done := compare(left.children[i], right.children[i]); done {
				return res, true
			}
		}
		if len(left.children) == len(right.children) {
			return false, false
		} else {
			return len(left.children) < len(right.children), true
		}
	}
}
