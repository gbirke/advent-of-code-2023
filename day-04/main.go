package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gbirke/advent-of-code-2023/pkg"
)

type ScratchCard struct {
	winningNumbers   []int
	availableNumbers []int
}

func parseNumberList(str string) []int {
	parts := strings.Split(str, " ")
	numbers := make([]int, 0, len(parts))
	for i, _ := range parts {
		if parts[i] == "" {
			continue
		}
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func ParseLine(line string) ScratchCard {
	cardAndNumbers := strings.Split(line, ":")
	parts := strings.Split(cardAndNumbers[1], " | ")

	return ScratchCard{parseNumberList(parts[0]), parseNumberList(parts[1])}
}

func countMatches(card ScratchCard) int {
	count := 0
	hash := make(map[int]bool)
	for _, wn := range card.winningNumbers {
		hash[wn] = true
	}
	for _, an := range card.availableNumbers {
		if hash[an] {
			count++
		}
	}
	return count
}

func main() {
	lines, err := pkg.ReadPuzzle()
	if err != nil {
		panic(err)
	}
	var points int
	for i, _ := range lines {
		card := ParseLine(lines[i])
		numMatches := countMatches(card)
		if numMatches > 0 {
			points += int(math.Pow(2, float64(numMatches-1)))
		}
	}
	fmt.Println("Part 1")
	fmt.Println(points)

}
