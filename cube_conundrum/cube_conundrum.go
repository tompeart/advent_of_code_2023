package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func isPossibleTurn(red, green, blue int) bool {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	return red <= maxRed && green <= maxGreen && blue <= maxBlue
}

func updateNumbers(colour, minColour, number *int) {
	if *number == 0 {
		return
	}
	*colour = *number
	if *colour > *minColour {
		*minColour = *number
	}
	*number = 0
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gameId := 0
	possibleSum := 0
	powerSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameId++
		var red, minRed int
		var green, minGreen int
		var blue, minBlue int
		number := 0
		possible := true
		for _, char := range scanner.Text() {
			switch char {
			case 'r':
				updateNumbers(&red, &minRed, &number)
			case 'g':
				updateNumbers(&green, &minGreen, &number)
			case 'b':
				updateNumbers(&blue, &minBlue, &number)
			case ';':
				if possible && !isPossibleTurn(red, green, blue) {
					possible = false
				}
				red = 0
				green = 0
				blue = 0
			}

			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(char))
				number = number*10 + digit
			} else if char != ' ' {
				number = 0
			}
		}

		if possible && isPossibleTurn(red, green, blue) {
			possibleSum += gameId
		}
		power := minRed * minGreen * minBlue
		powerSum += power
	}
	fmt.Printf("Sum of possible games: %v\n", possibleSum)
	fmt.Printf("Sum of powers: %v\n", powerSum)
}
