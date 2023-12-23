package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Schematic struct {
	xMax    int
	yMax    int
	numbers []Number
	symbols []Symbol
}

type Number struct {
	value  int
	xStart int
	xEnd   int
	y      int
}

type Symbol struct {
	x int
	y int
}

var validSymbols map[rune]struct{} = map[rune]struct{}{
	'*': {},
	'+': {},
	'$': {},
	'#': {},
}

func loadSchematic(inputFile string) Schematic {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []Number{}
	symbols := []Symbol{}
	xMax := 0
	y := 0
	for scanner.Scan() {
		currentValue := 0
		xStart := 0
		xEnd := 0
		for x, char := range scanner.Text() {
			if y == 0 {
				xMax++
			}
			if value, err := strconv.Atoi(string(char)); err == nil {
				if currentValue == 0 {
					xStart = x
				}
				xEnd = x
				currentValue = currentValue*10 + value
			} else {
				if currentValue != 0 {
					numbers = append(numbers, Number{
						value:  currentValue,
						xStart: xStart,
						xEnd:   xEnd,
						y:      y,
					})
					currentValue = 0
					xStart = 0
					xEnd = 0
				}
				if char != '.' {
					symbols = append(symbols, Symbol{
						x: x,
						y: y,
					})
				}
			}
		}
		if currentValue != 0 {
			numbers = append(numbers, Number{
				value:  currentValue,
				xStart: xStart,
				xEnd:   xEnd,
				y:      y,
			})
		}
		y++
	}
	return Schematic{
		xMax:    xMax - 1,
		yMax:    y - 1,
		numbers: numbers,
		symbols: symbols,
	}
}

func processSchematic(schematic Schematic) int {
	value := 0
	for _, number := range schematic.numbers {
		for y := -1; y < 2; y++ {
			if isSymbolInLocation(schematic, number.xStart-1, number.y+y) {
				value += number.value
				fmt.Printf("Found symbol for %v at %v, %v\n", number.value, number.xStart-1, number.y+y)
				break
			}
			if isSymbolInLocation(schematic, number.xEnd+1, number.y+y) {
				value += number.value
				fmt.Printf("Found symbol for %v at %v, %v\n", number.value, number.xEnd+1, number.y+y)
				break
			}
			if y != 0 {
				found := false
				for x := number.xStart; x <= number.xEnd; x++ {
					if isSymbolInLocation(schematic, x, number.y+y) {
						found = true
						fmt.Printf("Found symbol for %v at %v, %v\n", number.value, x, number.y+y)
						break
					}
				}
				if found {
					value += number.value
				}
			}
		}
	}
	return value
}

func isSymbolInLocation(schematic Schematic, x, y int) bool {
	if x < 0 || x > schematic.xMax || y < 0 || y > schematic.yMax {
		return false
	}
	for _, symbol := range schematic.symbols {
		if x == symbol.x && y == symbol.y {
			return true
		}
	}
	return false
}

func main() {
	schematic := loadSchematic("input.txt")
	for _, number := range schematic.numbers {
		fmt.Printf("Number %v, row %v, xStart %v, xEnd %v\n", number.value, number.y, number.xStart, number.xEnd)
	}
	for _, symbol := range schematic.symbols {
		fmt.Printf("Symbol at cell %v\n", symbol)
	}
	fmt.Printf("Schematic xMax %v, yMax %v\n", schematic.xMax, schematic.yMax)

	value := processSchematic(schematic)
	fmt.Printf("Schematic value is %v\n", value)
}
