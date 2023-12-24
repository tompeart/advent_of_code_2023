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

func calculateValue(schematic Schematic) int {
	value := 0
	for _, number := range schematic.numbers {
		for _, symbol := range schematic.symbols {
			if checkDiagonal(number, symbol) {
				value += number.value
				break
			}
		}
	}
	return value
}

func calculateGearRationSum(schematic Schematic) int {
	gearRatioSum := 0
	for _, symbol := range schematic.symbols {
		gearRatio := 0
		for _, number := range schematic.numbers {
			if checkDiagonal(number, symbol) {
				if gearRatio == 0 {
					gearRatio = number.value
				} else {
					gearRatio = gearRatio * number.value
					gearRatioSum += gearRatio
					break
				}
			}
		}
	}
	return gearRatioSum
}

func checkDiagonal(number Number, symbol Symbol) bool {
	return ((number.y-1 <= symbol.y && symbol.y <= number.y+1) &&
		(number.xStart-1 <= symbol.x && symbol.x <= number.xEnd+1))
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

	value := calculateValue(schematic)
	fmt.Printf("Schematic value is %v\n", value)

	gearRatioSum := calculateGearRationSum(schematic)
	fmt.Printf("Gear ratio sum is %v\n", gearRatioSum)
}
