package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func playCard(game string) int {
	setup := true
	matches := 0
	scoreCard := make(map[int]bool)
	current := 0
	for _, char := range game {
		if setup {
			switch char {
			case ' ':
				if current > 0 {
					scoreCard[current] = true
					current = 0
				}
			case '|':
				setup = false
			case ':':
				current = 0
			}
			if unicode.IsDigit(char) {
				current = processDigit(current, char)
			}
		} else {
			switch char {
			case ' ':
				if current > 0 {
					if scoreCard[current] {
						matches++
					}
					current = 0
				}
			case '|':
				setup = false
			case ':':
				current = 0
			}
			if unicode.IsDigit(char) {
				current = processDigit(current, char)
			}
		}
	}
	if current > 0 {
		if scoreCard[current] {
			matches++
		}
		current = 0
	}
	return matches
}

func processDigit(current int, digit rune) int {
	value, err := strconv.Atoi(string(digit))
	if err != nil {
		fmt.Printf("%v\n", string(digit))
		panic(err)
	}
	return current*10 + value
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	partOneScore := 0
	partTwoScore := 0
	copies := []int{}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i >= len(copies) {
			copies = append(copies, 1)
		} else {
			copies[i]++
		}
		matches := playCard(scanner.Text())

		if matches > 0 {
			partOneScore += (1 << (matches - 1))
			for j := i + 1; j <= i+matches; j++ {
				if j >= len(copies) {
					copies = append(copies, copies[i])
				} else {
					copies[j] += copies[i]
				}
			}
		}
		i++
	}

	for i := 0; i < len(copies); i++ {
		partTwoScore += copies[i]
	}
	fmt.Printf("Part one score: %v\n", partOneScore)
	fmt.Printf("Part two score: %v\n", partTwoScore)
}
