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
		var number int
		for _, char := range scanner.Text() {
			switch char {
			case 'r':
				if number > minRed {
					minRed = number
				}
				number = 0
			case 'g':
				if number > minGreen {
					minGreen = number
				}
				number = 0
			case 'b':
				if number > minBlue {
					minBlue = number
				}
				number = 0
			}

			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(char))
				number = number*10 + digit
			} else if char != ' ' {
				number = 0
			}
		}
		// calculate and add power
		power := minRed * minGreen * minBlue
		sum += power
		fmt.Printf("Game %v has power %v, sum = %v\n", gameId, power, sum)
	}
	fmt.Printf("Sum of powers: %v\n", sum)
}
