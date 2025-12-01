package main

import (
	"reflect"
	"testing"

	"github.com/djlechuck/advent-of-code/utils"
)

var (
	parsedSmallGrid = grid{
		robot: robot{coord{2, 2}},
		boxes: []box{
			{coord{3, 1}},
			{coord{5, 1}},
			{coord{4, 2}},
			{coord{4, 3}},
			{coord{4, 4}},
			{coord{4, 5}},
		},
		walls: []wall{
			{coord{0, 0}},
			{coord{1, 0}},
			{coord{2, 0}},
			{coord{3, 0}},
			{coord{4, 0}},
			{coord{5, 0}},
			{coord{6, 0}},
			{coord{7, 0}},
			{coord{0, 1}},
			{coord{7, 1}},
			{coord{0, 2}},
			{coord{1, 2}},
			{coord{7, 2}},
			{coord{0, 3}},
			{coord{7, 3}},
			{coord{0, 4}},
			{coord{2, 4}},
			{coord{7, 4}},
			{coord{0, 5}},
			{coord{7, 5}},
			{coord{0, 6}},
			{coord{7, 6}},
			{coord{0, 7}},
			{coord{1, 7}},
			{coord{2, 7}},
			{coord{3, 7}},
			{coord{4, 7}},
			{coord{5, 7}},
			{coord{6, 7}},
			{coord{7, 7}},
		},
	}
	parsedSmallMoves = moves{
		"<", "^", "^", ">", ">", ">", "v", "v", "<", "v", ">", ">", "v", "<", "<",
	}
)

func TestProcessInput(t *testing.T) {
	in := utils.ParseInput("../../../inputs/2024/15-2_example.txt")
	g, m := processInput(in)

	if !reflect.DeepEqual(g, parsedSmallGrid) {
		t.Errorf("processInput() - g: got %v, want %v", g, parsedSmallGrid)
	}
	if !reflect.DeepEqual(m, parsedSmallMoves) {
		t.Errorf("processInput() - m: got %v, want %v", m, parsedSmallMoves)
	}
}

func TestPartOne(t *testing.T) {
	expected := 2028

	result, _ := partOne(&parsedSmallGrid, parsedSmallMoves)
	if result != expected {
		t.Errorf("partOne() = %d, want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 0

	result, _ := partTwo()
	if result != expected {
		t.Errorf("partTwo() = %d, want %d", result, expected)
	}
}
