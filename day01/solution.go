package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var prefixToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func PartOne(input string) (answer int) {
	for _, line := range strings.Split(input, "\n") {
		var digits []int
		for _, r := range line {
			if r >= '1' && r <= '9' {
				digits = append(digits, int(r-'0'))
			}
		}
		answer += digits[0]*10 + digits[len(digits)-1]
	}

	return
}

func PartTwo(input string) (answer int) {
	for _, line := range strings.Split(input, "\n") {
		var digits []int
		for i, r := range line {
			if r >= '1' && r <= '9' {
				digits = append(digits, int(r-'0'))
			} else {
				for prefix, num := range prefixToDigit {
					if strings.HasPrefix(line[i:], prefix) {
						digits = append(digits, num)
					}
				}
			}
		}
		answer += digits[0]*10 + digits[len(digits)-1]
	}

	return
}

func main() {
	data, err := os.ReadFile("./day01/input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(input string) { defer wg.Done(); fmt.Println("Part 1", PartOne(input)) }(input)
	go func(input string) { defer wg.Done(); fmt.Println("Part 2", PartTwo(input)) }(input)

	wg.Wait()
}
