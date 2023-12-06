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

	var idSum, powerSum int

	// Open the txt file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {

		line := scanner.Text()

		// Split the line using whitespace, lines = games
		valid, id, power := checkGame(strings.Fields(line))

		if valid {
			idSum += id
		}
		powerSum += power

	}

	fmt.Println("Sum:", idSum)
	fmt.Println("Sum:", powerSum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

func checkGame(input []string) (bool, int, int) {

	colorMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	colorBuffer := map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}
	var err error
	var digitBuffer, idBuffer int

	powerBuffer := 1
	valid := true

	for _, field := range input {

		switch {

		case strings.Contains(field, "Game"):
			continue

		case strings.HasSuffix(field, ":"):
			idBuffer, err = strconv.Atoi(strings.SplitN(field, ":", 2)[0])
			valid = true
			continue

		case unicode.IsDigit([]rune(field)[0]):
			digitBuffer, err = strconv.Atoi(field)
			continue

		default:
			for key := range colorMap {
				if strings.Contains(field, key) {
					if digitBuffer > colorMap[key] {
						valid = false
					}
					if digitBuffer > colorBuffer[key] {
						colorBuffer[key] = digitBuffer
					}
					break
				}
			}
		}

		if err != nil {
			fmt.Println("Conversion error:", err)
			return false, 0, 0
		}
	}
	for _, value := range colorBuffer {
		powerBuffer *= value
	}
	return valid, idBuffer, powerBuffer

}
