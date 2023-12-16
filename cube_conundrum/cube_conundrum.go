package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gameId := 0
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameId++
		minRed := 0
		minGreen := 0
		minBlue := 0
		var lastNumber int
		var checkNextChar bool
		for _, char := range scanner.Text() {
			if char == ' ' {
				// fmt.Printf("Found space, last number: %v\n", lastNumber)
				checkNextChar = true
				continue
			}
			if unicode.IsDigit(char) {
				// fmt.Printf("Found digit %v, lastNumber %v", string(char), lastNumber)
				digit, _ := strconv.Atoi(string(char))
				lastNumber = lastNumber*10 + digit
				// fmt.Printf("-> %v\n", lastNumber)
				continue
			}

			if checkNextChar {
				switch char {
				case 'r':
					if lastNumber > minRed {
						minRed = lastNumber
					}
				case 'g':
					if lastNumber > minGreen {
						minGreen = lastNumber
					}
				case 'b':
					if lastNumber > minBlue {
						minBlue = lastNumber
					}
				}
				lastNumber = 0
			}
			checkNextChar = false
		}

		// calculate and add power
		power := minRed * minGreen * minBlue
		sum += power
		fmt.Printf("Game %v has power %v, sum = %v\n", gameId, power, sum)
	}
	fmt.Printf("Sum of powers: %v\n", sum)
}
