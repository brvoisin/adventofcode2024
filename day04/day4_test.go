// Advent of code 2024
// Day 4: Ceres Search
// https://adventofcode.com/2024/day/4
// Let's switch from Python to Go!

package main

import (
	"reflect"
	"testing"
)

func TestColumsToLines(t *testing.T) {
	input := []string{
		"ABC",
		"DEF",
		"GHI",
	}
	want := []string{
		"ADG",
		"BEH",
		"CFI",
	}
	got := ColumnsToLines(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ColumnsToLines() = %v, want %v", got, want)
	}
}

func TestDiagonal1ToLines(t *testing.T) {
	arg := []string{
		"ABC",
		"DEF",
		"GHI",
	}
	want := []string{
		"A",
		"BD",
		"CEG",
		"FH",
		"I",
	}

	got := Diagonal1ToLines(arg)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diagonal1ToLines() = %v, want %v", got, want)
	}
}

func TestDiagonal2ToLines(t *testing.T) {
	arg := []string{
		"ABC",
		"DEF",
		"GHI",
	}
	want := []string{
		"C",
		"BF",
		"AEI",
		"DH",
		"G",
	}

	got := Diagonal2ToLines(arg)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diagonal2ToLines() = %v, want %v", got, want)
	}
}

func TestCountAllWords(t *testing.T) {
	sample := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	want := 18

	got := CountAllWords(sample)

	if got != want {
		t.Errorf("CountAllWords() = %v, want %v", got, want)
	}
}
