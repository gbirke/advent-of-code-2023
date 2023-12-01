package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gbirke/advent-of-code-2023/pkg"
)

type FirstAndLastMatch struct {
	firstMatch string
	lastMatch  string
}

func GetFirstAndLastMatchFromLine(line string, numberwords []string) (FirstAndLastMatch, error) {
	firstNumber := ""
	lastNumber := ""
	firstIndex := len(line) + 1
	lastIndex := -1

	for _, word := range numberwords {
		wordIndex := strings.Index(line, word)
		lastWordIndex := strings.LastIndex(line, word)
		if wordIndex > -1 && wordIndex < firstIndex {
			firstIndex = wordIndex
			firstNumber = word
		}
		if lastWordIndex > -1 && lastWordIndex > lastIndex {
			lastIndex = lastWordIndex
			lastNumber = word
		}
	}

	if firstIndex == len(line) + 1 || lastIndex == -1 {
		return FirstAndLastMatch{"",""}, fmt.Errorf("No number words found in line %q", line)
	}

	return FirstAndLastMatch{firstNumber, lastNumber}, nil
}

func GetNumberFromLineWithSimpleWordList(line string) (int, error) {
	numberwords := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	match, err := GetFirstAndLastMatchFromLine(line, numberwords)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(match.firstMatch + match.lastMatch)
}


type getNumberFromLine func(string) (int, error)

func GetChecksum(lines []string, fn getNumberFromLine) int {
	checksum := 0
	for _, line := range lines {
		number, _ := fn(line)
		checksum += number
	}
	return checksum
}

func main() {
	lines, err := pkg.ReadPuzzle()
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1")
	fmt.Println("checksum is", GetChecksum(lines, GetNumberFromLineWithSimpleWordList))
}
