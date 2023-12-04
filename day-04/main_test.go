package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	got := ParseLine("Card 1: 1 21 53 59 44 | 69 82 63 72 16 21 14  1")
	want := ScratchCard{[]int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseLine() = %v", got)
	}
}

func TestCountMatches(t *testing.T) {
	var tests = []struct {
		input ScratchCard
		want  int
	}{
		{ScratchCard{[]int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}}, 4},
		{ScratchCard{[]int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}}, 2},
		{ScratchCard{[]int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}}, 2},
		{ScratchCard{[]int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}}, 1},
		{ScratchCard{[]int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}}, 0},
		{ScratchCard{[]int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}}, 0},
	}

	for _, test := range tests {
		if got := countMatches(test.input); got != test.want {
			t.Errorf("countMatches(%v) = %v", test.input, got)
		}
	}
}
