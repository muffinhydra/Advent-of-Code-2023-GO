package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Scaffold struct {
	data [][]int
}

func main() {

	// Open the txt file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fields [][]int
	// Read line by line
	for scanner.Scan() {

		line := scanner.Text()
		var digits []int
		for _, digit := range strings.Fields(line) {
			t, _ := strconv.Atoi(digit)
			digits = append(digits, t)
		}
		fields = append(fields, digits)

	}

	scaffolds := createScaffolds(fields)

	predictions := createPredictions(scaffolds)

	fmt.Println("sum part 1:", addPredictions(predictions))

	pastNumbers := createReversePredictions(scaffolds)

	fmt.Println("sum part 2:", addPart2Predictions(pastNumbers))
}

func buildScaffold(data []int) []int {
	var slice []int
	for i := 1; i < len(data); i++ {
		slice = append(slice, data[i]-data[i-1])

	}
	if len(slice) == 0 {
		return []int{0}
	}

	return slice
}

func createScaffolds(fields [][]int) []Scaffold {
	var result []Scaffold
	for _, line := range fields {
		var subSliceCollection [][]int
		subSliceCollection = append(subSliceCollection, line)
		subSlice := line
		for {

			subSlice = buildScaffold(subSlice)

			subSliceCollection = append(subSliceCollection, subSlice)

			if slices.Max(subSlice) == 0 && slices.Min(subSlice) == 0 {
				break
			}

		}
		var t Scaffold
		slices.Reverse(subSliceCollection)
		t.data = subSliceCollection
		result = append(result, t)

	}
	return result
}

func createPredictions(scaffolds []Scaffold) []Scaffold {

	for _, scaffold := range scaffolds {

		var number int
		for i, slice := range scaffold.data {
			number += slice[len(slice)-1]
			slice = append(slice, number)
			scaffold.data[i] = slice

		}

	}
	return scaffolds
}

func createReversePredictions(scaffolds []Scaffold) []Scaffold {
	for _, scaffold := range scaffolds {
		fmt.Println("Processing scaffold")

		var number int
		for i := 1; i < len(scaffold.data); i++ {

			fmt.Printf("Processing slice %d\n", i)

			number = scaffold.data[i][0] - scaffold.data[i-1][0]

			fmt.Println("raw slice:", scaffold.data[i])

			scaffold.data[i] = append([]int{number}, scaffold.data[i]...)

			fmt.Println("Updated slice:", scaffold.data[i])

		}
	}
	fmt.Println("Finished processing")
	return scaffolds
}

func addPredictions(scaffolds []Scaffold) int {
	var sum int
	for _, scaffold := range scaffolds {
		lastNumber := len(scaffold.data) - 1
		sum += scaffold.data[lastNumber][len(scaffold.data[lastNumber])-1]
	}
	return sum
}
func addPart2Predictions(scaffolds []Scaffold) int {
	var sum int
	for _, scaffold := range scaffolds {
		lastNumber := len(scaffold.data) - 1
		sum += scaffold.data[lastNumber][0]
	}
	return sum
}
