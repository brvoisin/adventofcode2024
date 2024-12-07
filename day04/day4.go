// Advent of code 2024
// Day 4: Ceres Search
// https://adventofcode.com/2024/day/4
// Let's switch from Python to Go!
//
// Usage:
//   go run day4.go < input

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const word = "XMAS"
const reverseWord = "SAMX"

func main() {
	fmt.Println(CountAllWords(ReadLines(os.Stdin)))
}

func ReadLines(file *os.File) (lines []string) {
	scanner := bufio.NewScanner(file)
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
