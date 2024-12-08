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

var sample = []string{
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

func TestCountAllWords(t *testing.T) {
	want := 18

	got := CountAllWords(sample)

	if got != want {
		t.Errorf("CountAllWords() = %v, want %v", got, want)
	}
}

func TestCountAllXMAS(t *testing.T) {
	want := 9

	got := CountAllXMAS(sample)

	if got != want {
		t.Errorf("CountAllXMAS() = %v, want %v", got, want)
	}
}
