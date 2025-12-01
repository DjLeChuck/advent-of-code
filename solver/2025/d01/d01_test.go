package main

import (
	"reflect"
	"testing"

	"github.com/djlechuck/advent-of-code/utils"
)

var (
	parsedInput = []instruction{
		{op: "L", length: 68},
		{op: "L", length: 30},
		{op: "R", length: 48},
		{op: "L", length: 5},
		{op: "R", length: 60},
		{op: "L", length: 55},
		{op: "L", length: 1},
		{op: "L", length: 99},
		{op: "R", length: 14},
		{op: "L", length: 82},
	}
)

func TestProcessInput(t *testing.T) {
	in := utils.ParseInput("../../../inputs/2025/01_example.txt")
	e := processInput(in)

	if !reflect.DeepEqual(e, parsedInput) {
		t.Errorf("processInput(): got %v, want %v", e, parsedInput)
	}
}

func TestPartOne(t *testing.T) {
	expected := 3

	result, _ := partOne(parsedInput)
	if result != expected {
		t.Errorf("partOne() = %d, want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 6

	result, _ := partTwo(parsedInput)
	if result != expected {
		t.Errorf("partTwo() = %d, want %d", result, expected)
	}
}
