package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Open the txt file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum, sumDouble int
	lineCount := 0
	cardCopies := make(map[int]int)
	// Read line by line
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		doubleIt := 0
		cleanedLine := strings.Split(line, ":")
		splitLine := strings.Split(cleanedLine[1], "|")
		winnings := strings.Fields(splitLine[0])
		scratched := strings.Fields(splitLine[1])

		futureCard := 0
		for _, num := range scratched {
			for _, winner := range winnings {
				if num == winner {
					//part 1
					if doubleIt != 0 {
						doubleIt *= 2
					} else {
						doubleIt = 1
					}
					//part 2
					futureCard++
					cardCopies[lineCount+futureCard] += 1 + cardCopies[lineCount]

				}

			}
		}
		sumDouble += doubleIt
	}

	for _, copies := range cardCopies {
		sum += copies
	}

	fmt.Println("SumDouble :", sumDouble)
	fmt.Println("Sum :", sum+lineCount)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
