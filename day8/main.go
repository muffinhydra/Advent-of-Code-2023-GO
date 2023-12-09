package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	input       string
	outputRight string
	outputLeft  string
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

	var nodes []Node
	var directions []string
	// Read line by line
	for scanner.Scan() {

		line := scanner.Text()
		splits := strings.Fields(line)
		if len(splits) == 0 {
			continue
		}
		if len(splits) == 1 {
			directions = strings.Split(splits[0], "")
			continue
		}
		var node Node

		outputLeft := strings.Split(splits[2], ",")[0]
		outputLeft = strings.Split(outputLeft, "(")[1]
		outputRight := strings.Split(splits[3], ")")[0]

		node.input = splits[0]
		node.outputLeft = outputLeft
		node.outputRight = outputRight
		nodes = append(nodes, node)

	}

	var startingPoints []string
	var lengths []int64

	for _, node := range nodes {
		if strings.HasSuffix(node.input, "A") {
			startingPoints = append(startingPoints, node.input)
		}
	}
	var i int64
	for pos := range startingPoints {

	pathfinding:
		for {
			for _, direction := range directions {
				for _, node := range nodes {

					if node.input == startingPoints[pos] {
						if direction == "R" {
							startingPoints[pos] = node.outputRight
							break
						}
						if direction == "L" {
							startingPoints[pos] = node.outputLeft
							break
						}
					}
				}

				i++

				if strings.HasSuffix(startingPoints[pos], "Z") {
					lengths = append(lengths, i)
					i = 0
					break pathfinding
				}

			}

		}
	}
	fmt.Println("startingPoints :", startingPoints)
	fmt.Println("Lengths :", lengths)
	fmt.Println("Steps :", calculateLCM(lengths))

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

func calculateGCD(a, b int64) int64 {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func calculateLCM(numbers []int64) int64 {
    lcm := numbers[0]
    for i := 1; i < len(numbers); i++ {
        gcd := calculateGCD(lcm, numbers[i])
        lcm = lcm * numbers[i] / gcd
    }
    return lcm
}

