package main

import (
	"bufio"
	"fmt"
	"math"
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

	var scaffolds []Scaffold
	scaffolds := createScaffolds(fields)
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
		scaffolds = append(scaffolds, t)

	}

	for _, scaffold := range scaffolds {
		fmt.Println("scaffold :", scaffold)
		var number int
		for i, slice := range scaffold.data {
			fmt.Println("slice :", slice)
			number += slice[len(slice)-1]
			slice = append(slice, number)
			scaffold.data[i] = slice

		}

	}

	var sum int
	for _, scaffold := range scaffolds {
		lastNumber := len(scaffold.data) - 1
		sum += scaffold.data[lastNumber][len(scaffold.data[lastNumber])-1]
	}
	fmt.Println("sum :", sum)
}

func buildScaffold(data []int) []int {
	var slice []int
	for i := 1; i < len(data); i++ {
		slice = append(slice, data[i] - data[i-1])
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