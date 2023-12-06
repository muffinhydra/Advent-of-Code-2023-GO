package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	numbersMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		fmt.Println(scanner.Text())

		lineLeft := replaceLeftMostMatch(scanner.Text(), numbersMap)
		lineRight := replaceRightMostMatch(scanner.Text(), numbersMap)

		var bufferOne, bufferTwo string
		for _, char := range lineLeft {
			if unicode.IsDigit(char) {
				bufferOne = string(char)
				break
			}
		}
		reversedString := reverseString(lineRight)
		for _, char := range reversedString {
			if unicode.IsDigit(char) {
				bufferTwo = string(char)
				break
			}
		}
		concatNumber, err := strconv.Atoi(bufferOne + bufferTwo)
		if err != nil {
			fmt.Println("Conversion error:", err)
			return
		}
		sum += concatNumber
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Sum:", sum)
}

func reverseString(input string) string {
	runes := []rune(input)

	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func replaceLeftMostMatch(input string, myMap map[string]int) string {
	lowestIndex := len(input)
	match := ""

	for key := range myMap {
		index := strings.Index(input, key)
		if index != -1 && index < lowestIndex {
			lowestIndex = index
			match = key
		}
	}

	if match != "" {
		digit := fmt.Sprintf("%d", myMap[match])
		return input[:lowestIndex] + digit + input[lowestIndex+len(match):]
	}
	return input
}

func replaceRightMostMatch(input string, myMap map[string]int) string {
	highestIndex := -1
	match := ""

	for key := range myMap {
		index := strings.LastIndex(input, key)
		if index != -1 && index > highestIndex {
			highestIndex = index
			match = key
		}
	}

	if match != "" {
		digit := fmt.Sprintf("%d", myMap[match])
		return input[:highestIndex] + digit + input[highestIndex+len(match):]
	}
	return input
}
