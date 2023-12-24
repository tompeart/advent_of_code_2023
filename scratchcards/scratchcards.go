package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func scoreGame(game string) int {
	setup := true
	exponent := -1
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
						exponent++
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
			exponent++
		}
		current = 0
	}
	if exponent < 0 {
		return 0
	}
	return 1 << exponent
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

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		score += scoreGame(scanner.Text())
	}
	fmt.Printf("Sum of scores: %v\n", score)
}
