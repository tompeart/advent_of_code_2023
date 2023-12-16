package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func checkTurn(red, green, blue int) bool {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	// fmt.Printf("Checking red %v, green %v, blue %v\n", red, green, blue)
	if red <= maxRed && green <= maxGreen && blue <= maxBlue {
		// fmt.Println("Good")
		return true
	}
	// fmt.Println("Bad")
	return false
}

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
		red := 0
		green := 0
		blue := 0
		possible := true
		var lastNumber int
		var endOfTurn bool
		var checkNextChar bool
		for _, char := range scanner.Text() {
			if char == ' ' {
				// fmt.Printf("Found space, last number: %v\n", lastNumber)
				checkNextChar = true
				continue
			}
			if char == ';' {
				endOfTurn = true
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
					// fmt.Printf("R - %v\n", lastNumber)
					red = lastNumber
				case 'g':
					// fmt.Printf("G - %v\n", lastNumber)
					green = lastNumber
				case 'b':
					// fmt.Printf("B - %v\n", lastNumber)
					blue = lastNumber
				}
				lastNumber = 0
			}
			checkNextChar = false

			if endOfTurn {
				if !checkTurn(red, green, blue) {
					possible = false
					break
				}
				red = 0
				green = 0
				blue = 0
				endOfTurn = false
			}
		}
		// check final turn in line
		if possible && checkTurn(red, green, blue) {
			sum += gameId
			fmt.Printf("Game %v is possible, sum = %v\n", gameId, sum)
		}
	}
	fmt.Printf("Sum of possible games: %v\n", sum)
}
