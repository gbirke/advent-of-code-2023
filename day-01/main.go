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

	if firstIndex == len(line)+1 || lastIndex == -1 {
		return FirstAndLastMatch{"", ""}, fmt.Errorf("No number words found in line %q", line)
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

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

var wordMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func GetNumberFromLineWithMappedWordList(line string) (int, error) {

	match, err := GetFirstAndLastMatchFromLine(line, keys(wordMap))
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(wordMap[match.firstMatch] + wordMap[match.lastMatch])
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

	fmt.Println("Part 2")
	fmt.Println("checksum is", GetChecksum(lines, GetNumberFromLineWithMappedWordList))
}
