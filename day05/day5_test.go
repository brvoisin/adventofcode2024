package main

import (
	"reflect"
	"strings"
	"testing"
)

var sample = strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

var sampleRules = []PageOrderingRule{
	{Before: 47, After: 53},
	{Before: 97, After: 13},
	{Before: 97, After: 61},
	{Before: 97, After: 47},
	{Before: 75, After: 29},
	{Before: 61, After: 13},
	{Before: 75, After: 53},
	{Before: 29, After: 13},
	{Before: 97, After: 29},
	{Before: 53, After: 29},
	{Before: 61, After: 53},
	{Before: 97, After: 53},
	{Before: 61, After: 29},
	{Before: 47, After: 13},
	{Before: 75, After: 47},
	{Before: 97, After: 75},
	{Before: 47, After: 61},
	{Before: 75, After: 61},
	{Before: 47, After: 29},
	{Before: 75, After: 13},
	{Before: 53, After: 13},
}

var sampleUpdates = []Update{
	{75, 47, 61, 53, 29},
	{97, 61, 53, 29, 13},
	{75, 29, 13},
	{75, 97, 47, 61, 53},
	{61, 13, 29},
	{97, 13, 75, 29, 47},
}

var sampleInput = Input{Rules: sampleRules, Updates: sampleUpdates}

func TestParse(t *testing.T) {
	want := sampleInput

	got, err := Parse(sample)

	if err != nil {
		t.Errorf("Parse() error = %#v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %#v, want %#v", got, want)
	}
}

func TestIsCorrectlyOrdered(t *testing.T) {
	type args struct {
		rules  []PageOrderingRule
		update Update
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Correctly ordered",
			args: args{
				rules:  sampleRules,
				update: sampleUpdates[0],
			},
			want: true,
		},
		{
			name: "Incorrectly ordered",
			args: args{
				rules:  sampleRules,
				update: sampleUpdates[3],
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCorrectlyOrdered(tt.args.rules, tt.args.update); got != tt.want {
				t.Errorf("isCorrectlyOrdered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfMiddlePageNumberOfCorrectlyOrderedUpdates(t *testing.T) {
	want := 143

	got := SumOfMiddlePageNumberOfCorrectlyOrderedUpdates(sampleInput)

	if got != want {
		t.Errorf("SumOfMiddlePageNumberOfCorrectlyOrderedUpdates() = %v, want %v", got, want)
	}
}
