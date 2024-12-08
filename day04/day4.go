// Advent of code 2024
// Day 4: Ceres Search
// https://adventofcode.com/2024/day/4
// Let's switch from Python to Go!
//
// Usage:
//   go run day4.go part1 < input
//   go run day4.go part2 < input

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	word        = "XMAS"
	reverseWord = "SAMX"
	mas         = "MAS"
	sam         = "SAM"
)
const (
	part1 = "part1"
	part2 = "part2"
)

func main() {
	part := os.Args[1]
	var exerice func([]string) int
	switch part {
	case part1:
		exerice = CountAllWords
	case part2:
		exerice = CountAllXMAS
	default:
		fmt.Fprintf(os.Stderr, "part argument must be %s or %s\n", part1, part2)
		os.Exit(1)
	}
	fmt.Println(exerice(ReadLines(os.Stdin)))
}

func ReadLines(r io.Reader) (lines []string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func CountAllWords(input []string) (count int) {
	for _, dimension := range GetAllDimensions(input) {
		for _, line := range dimension {
			count += strings.Count(line, word) + strings.Count(line, reverseWord)
		}
	}
	return
}

func GetAllDimensions(input []string) [][]string {
	return [][]string{
		input,
		ColumnsToLines(input),
		Diagonal1ToLines(input),
		Diagonal2ToLines(input),
	}
}

func ColumnsToLines(input []string) []string {
	// Assuming input is not empty and all lines have the same length.
	width := len(input[0])
	height := len(input)
	var result = make([]string, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			result[j] += string(input[i][j])
		}
	}
	return result
}

func Diagonal1ToLines(input []string) []string {
	// Assuming input is not empty and all lines have the same length.
	width := len(input[0])
	height := len(input)
	var result = make([]string, height+width-1)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			result[x+y] += string(input[x][y])
		}
	}
	return result
}

func Diagonal2ToLines(input []string) []string {
	// Assuming input is not empty and all lines have the same length.
	width := len(input[0])
	height := len(input)
	var result = make([]string, height+width-1)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			result[x+y] += string(input[x][height-1-y])
		}
	}
	return result
}

func CountAllXMAS(input []string) (count int) {
	// Assuming input is not empty and all lines have the same length.
	width := len(input[0])
	height := len(input)
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			word1 := string(input[x-1][y-1]) + string(input[x][y]) + string(input[x+1][y+1])
			word2 := string(input[x-1][y+1]) + string(input[x][y]) + string(input[x+1][y-1])
			if (word1 == mas || word1 == sam) && (word2 == mas || word2 == sam) {
				count++
			}
		}
	}
	return
}
