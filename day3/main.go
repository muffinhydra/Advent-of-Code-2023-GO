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
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var field []string
	var sum int

	for scanner.Scan() {
		field = append(field, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Extract positions of numbers and their values from the field
	numPositions := extractNumbersAndPositions(field)

	// Find positions of all valid special symbols
	specialSymbolPositions := findSpecialSymbolPositions(field)

	// Find adjacent numbers to special symbols
	adjacentNumbers := findAdjacentNumbersToSpecial(numPositions, specialSymbolPositions)

	var gearPairs [][]int // Store pairs of adjacent numbers

	// Iterate through positions of '*' symbol
	for _, pos := range specialSymbolPositions['*'] {
		starPosition := make(map[rune][][]int)
		starPosition['*'] = [][]int{
			pos,
		}

		// Find adjacent numbers to the '*' symbol
		adjacentNumberBuffer := findAdjacentNumbersToSpecial(numPositions, starPosition)

		// If exactly two adjacent numbers are found, add them to gearPairs
		if len(adjacentNumberBuffer) == 2 {
			gearPairs = append(gearPairs, adjacentNumberBuffer)
		}
	}

	gearRatioSum := multiplyAndSumPairs(gearPairs)

	for _, num := range adjacentNumbers {
		sum += num
	}

	fmt.Println("Sum of adjacent numbers:", sum)
	fmt.Println("Gear Ratio Sum of * adjacent numbers:", gearRatioSum)
}

// Extracts numbers and their positions from the field
func extractNumbersAndPositions(field []string) map[string][][]int {
	numPositions := make(map[string][][]int)

	for r, row := range field {
		col := 0
		for col < len(row) {
			// Extracts continuous digits as a number
			if unicode.IsDigit(rune(row[col])) {
				numStr := ""
				startCol := col
				for col < len(row) && unicode.IsDigit(rune(row[col])) {
					numStr += string(row[col])
					col++
				}

				// Generate a unique key for the number
				key := numStr + ".1"
				for {
					if _, exists := numPositions[key]; !exists {
						break
					}
					increment, _ := strconv.Atoi(strings.Split(key, ".")[1])
					key = fmt.Sprintf("%s.%d", numStr, increment+1)
				}

				// Store the number's positions with the generated key
				numPositions[key] = append(numPositions[key], []int{r + 1, startCol + 1})
				for i := 1; i < len(numStr); i++ {
					numPositions[key] = append(numPositions[key], []int{r + 1, startCol + i + 1})
				}
			} else {
				col++
			}
		}
	}

	return numPositions
}

// Checks if a character is a valid symbol
func isValidSymbol(input rune) bool {
	if (input >= '0' && input <= '9') || input == '.' {
		return false
	}
	return true
}

// Finds positions of special symbols in the field
func findSpecialSymbolPositions(field []string) map[rune][][]int {
	specialSymbolPositions := make(map[rune][][]int)

	for r, row := range field {
		for c, char := range row {
			if isValidSymbol(char) {
				if _, ok := specialSymbolPositions[char]; !ok {
					specialSymbolPositions[char] = [][]int{{r + 1, c + 1}}
				} else {
					specialSymbolPositions[char] = append(specialSymbolPositions[char], []int{r + 1, c + 1})
				}
			}
		}
	}

	return specialSymbolPositions
}

// Finds adjacent numbers to special symbols
func findAdjacentNumbersToSpecial(numPositionsInMemory map[string][][]int, specialSymbolPositionsInMemory map[rune][][]int) []int {

	// Today I learned maps in GO are not copied but referenced as parameters
	// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHHHHHHHHHHHHHHHHHHHHHH
	// Copy from the original map to the target map

	numPositions := make(map[string][][]int)
	for key, value := range numPositionsInMemory {
		numPositions[key] = value
	}

	specialSymbolPositions := make(map[rune][][]int)
	for key, value := range specialSymbolPositionsInMemory {
		specialSymbolPositions[key] = value
	}

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	var adjacentNumbers []int

	// Iterate through positions of special symbols
	for _, positions := range specialSymbolPositions {
		for _, pos := range positions {
			r, c := pos[0], pos[1]

			// Check adjacent positions for numbers
			for _, dir := range directions {
				newR, newC := r+dir[0], c+dir[1]

				for key, numPos := range numPositions {
					for _, nPos := range numPos {
						nR, nC := nPos[0], nPos[1]
						if newR == nR && newC == nC {
							decodedNum := decodeFormattedKey(key)
							adjacentNumbers = append(adjacentNumbers, decodedNum)
							delete(numPositions, key)
						}
					}
				}
			}
		}
	}
	return adjacentNumbers
}

// Decodes the formatted key to extract the number
func decodeFormattedKey(key string) int {
	parts := strings.Split(key, ".")
	numberPart := parts[0]

	decodedNumber, err := strconv.Atoi(numberPart)
	if err != nil {
		fmt.Println("Error decoding key:", err)
		return 0
	}

	return decodedNumber
}

// Multiplies and sums the pairs of adjacent numbers
func multiplyAndSumPairs(adjacentNumbers [][]int) int {
	sum := 0
	for _, nums := range adjacentNumbers {
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				sum += nums[i] * nums[j]
			}
		}
	}

	return sum
}
