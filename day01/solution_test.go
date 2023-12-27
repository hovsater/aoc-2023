package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	expectedAnswer := 142

	t.Parallel()
	if answer := PartOne(input); answer != expectedAnswer {
		t.Errorf("PartOne(%q) = %d; want %d", input, answer, expectedAnswer)
	}
}

func TestPartTwo(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	expectedAnswer := 281

	t.Parallel()
	if answer := PartTwo(input); answer != expectedAnswer {
		t.Errorf("PartTwo(%q) = %d; want %d", input, answer, expectedAnswer)
	}
}
