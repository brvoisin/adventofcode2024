package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	part1 = "part1"
	part2 = "part2"
)

type PageOrderingRule struct {
	Before int
	After  int
}

type Update []int

type Input struct {
	Rules   []PageOrderingRule
	Updates []Update
}

func main() {
	part := os.Args[1]
	var exerice func(Input) int
	switch part {
	case part1:
		exerice = SumOfMiddlePageNumberOfCorrectlyOrderedUpdates
	case part2:
		exerice = SumOfMiddlePageNumberOfIncorrectlyOrderedUpdates
	default:
		fmt.Fprintf(os.Stderr, "part argument must be %s or %s\n", part1, part2)
		os.Exit(1)
	}
	input, err := Parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(exerice(input))
}

func Parse(r io.Reader) (Input, error) {
	input := Input{}
	var err error
	scanner := bufio.NewScanner(r)
	input.Rules, err = parseRules(scanner)
	if err != nil {
		return input, err
	}
	input.Updates, err = parseUpdates(scanner)
	if err != nil {
		return input, err
	}
	return input, nil
}

func parseRules(scanner *bufio.Scanner) ([]PageOrderingRule, error) {
	var rules []PageOrderingRule
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		before, err := strconv.Atoi(parts[0])
		if err != nil {
			return rules, err
		}
		after, err := strconv.Atoi(parts[1])
		if err != nil {
			return rules, err
		}
		rules = append(rules, PageOrderingRule{Before: before, After: after})
	}
	return rules, nil
}

func parseUpdates(scanner *bufio.Scanner) ([]Update, error) {
	var updates []Update
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		update := make(Update, len(parts))
		for i, part := range parts {
			value, err := strconv.Atoi(part)
			if err != nil {
				return updates, err
			}
			update[i] = value
		}
		updates = append(updates, update)
	}
	return updates, nil
}

func SumOfMiddlePageNumberOfCorrectlyOrderedUpdates(input Input) (result int) {
	return sumOfMiddlePageNumber(getCorrectlyOrderedUpdates(input))
}

func sumOfMiddlePageNumber(updates []Update) (result int) {
	for _, update := range updates {
		result += update[len(update)/2]
	}
	return
}

func getCorrectlyOrderedUpdates(input Input) (result []Update) {
	for _, update := range input.Updates {
		if isCorrectlyOrdered(input.Rules, update) {
			result = append(result, update)
		}
	}
	return
}

func isCorrectlyOrdered(rules []PageOrderingRule, update Update) bool {
	for _, rule := range rules {
		indexBefore := slices.Index(update, rule.Before)
		indexAfter := slices.Index(update, rule.After)
		if indexBefore == -1 || indexAfter == -1 {
			continue
		}
		if indexBefore > indexAfter {
			return false
		}
	}
	return true
}

func SumOfMiddlePageNumberOfIncorrectlyOrderedUpdates(input Input) (result int) {
	return sumOfMiddlePageNumber(putAllUpdatesInTheRightOrder(input.Rules, getIncorrectlyOrderedUpdates(input)))
}

func getIncorrectlyOrderedUpdates(input Input) (result []Update) {
	for _, update := range input.Updates {
		if !isCorrectlyOrdered(input.Rules, update) {
			result = append(result, update)
		}
	}
	return
}

func putAllUpdatesInTheRightOrder(rules []PageOrderingRule, updates []Update) []Update {
	var result = make([]Update, len(updates))
	for i, update := range updates {
		result[i] = putUpdateInTheRightOrder(rules, update)
	}
	return result
}

func putUpdateInTheRightOrder(rules []PageOrderingRule, update Update) Update {
	result := slices.Clone(update)
	slices.SortFunc(result, func(i, j int) int {
		for _, rule := range rules {
			if i == rule.Before && j == rule.After {
				return -1
			}
			if i == rule.After && j == rule.Before {
				return 1
			}
		}
		return 0
	})
	return result
}
