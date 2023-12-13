package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var digitsToNumbers = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

type Node struct {
	children   map[rune]*Node
	endOfDigit int
}

type Trie struct {
	root *Node
}

func loadDigitIntoTrie(trie *Trie, word string, digit int) {
	current := trie.root
	for _, char := range word {
		if _, ok := current.children[char]; !ok {
			current.children[char] = &Node{
				children: make(map[rune]*Node),
			}
		}
		current = current.children[char]
	}
	current.endOfDigit = digit
}

func printWordsInTrie(node *Node, prefix string) {
	for char, nextNode := range node.children {
		printWordsInTrie(nextNode, prefix+string(char))
	}
	if node.endOfDigit != 0 {
		fmt.Println(prefix + " -> " + strconv.Itoa(node.endOfDigit))
	}
}

func findDigitInLine(trie *Trie, line string) int {
	current := trie.root
	for _, char := range line {
		if next, ok := current.children[char]; ok {
			current = next
			if current.endOfDigit != 0 {
				return current.endOfDigit
			}
		} else {
			return 0
		}
	}
	return 0
}

func main() {
	numbers := &Trie{
		root: &Node{
			children: make(map[rune]*Node),
		},
	}
	for digit, number := range digitsToNumbers {
		loadDigitIntoTrie(numbers, strconv.Itoa(digit), digit)
		loadDigitIntoTrie(numbers, number, digit)
	}

	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int
	for scanner.Scan() {
		line := scanner.Text()
		var first int
		var last int
		for i := range line {
			foundDigit := findDigitInLine(numbers, line[i:])
			if foundDigit == 0 {
				continue
			}
			if first == 0 {
				first = foundDigit
			}
			last = foundDigit
		}
		calibration := first*10 + last
		total += calibration
	}
	fmt.Printf("Sum of calibration values: %d\n", total)
}
