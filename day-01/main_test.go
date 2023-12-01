package main

import (
	"testing"
)

func TestGetNumberFromLineWithSimpleWordList(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		if got, _ := GetNumberFromLineWithSimpleWordList(test.input); got != test.want {
			t.Errorf("GetNumberFromLine(%q) = %v", test.input, got)
		}
	}
}

func TestGetNumberFromLineWithMappedWordList(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"eighttwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneeigt234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, test := range tests {
		if got, _ := GetNumberFromLineWithMappedWordList(test.input); got != test.want {
			t.Errorf("GetNumberFromLine(%q) = %v", test.input, got)
		}
	}
}

func TestGetChecksum(t *testing.T) {
	got := GetChecksum([]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}, GetNumberFromLineWithSimpleWordList)
	want := 142

	if got != want {
		t.Errorf("GetChecksum() = %v", got)
	}
}
