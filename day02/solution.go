package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var colorRegexp = regexp.MustCompile(`(\d+) (red|green|blue)`)

func PartOne(input string) (answer int) {
	for i, line := range strings.Split(input, "\n") {
		max := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, matches := range colorRegexp.FindAllStringSubmatch(line, -1) {
			if num, err := strconv.Atoi(matches[1]); err == nil && num > max[matches[2]] {
				max[matches[2]] = num
			}
		}

		if max["red"] <= 12 && max["green"] <= 13 && max["blue"] <= 14 {
			answer += i + 1
		}
	}

	return
}

func PartTwo(input string) (answer int) {
	for _, line := range strings.Split(input, "\n") {
		max := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, matches := range colorRegexp.FindAllStringSubmatch(line, -1) {
			if num, err := strconv.Atoi(matches[1]); err == nil && num > max[matches[2]] {
				max[matches[2]] = num
			}
		}

		answer += max["red"] * max["green"] * max["blue"]
	}

	return
}

func main() {
	data, err := os.ReadFile("./day02/input.txt")
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
