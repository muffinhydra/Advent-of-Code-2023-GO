package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

	dataMap := make(map[string][]int)
	var part2Cache []int

	// Read line by line
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if line[0] == "Time:" {
			var time string

			for i := 1; i < len(line); i++ {
				num, _ := strconv.Atoi(line[i])
				dataMap["time"] = append(dataMap["time"], num)
				time = time + line[i]
			}

			num, _ := strconv.Atoi(time)
			part2Cache = append(part2Cache, num)
		}
		if line[0] == "Distance:" {
			var distance string
			for i := 1; i < len(line); i++ {
				num, _ := strconv.Atoi(line[i])
				dataMap["distance"] = append(dataMap["distance"], num)
				distance = distance + line[i]
			}

			num, _ := strconv.Atoi(distance)
			part2Cache = append(part2Cache, num)
		}
	}

	var cache []int

	for i := 0; i < len(dataMap["time"]); i++ {
		cache = append(cache, len(findPossibleHoldTimesBruteForce(dataMap["time"][i], dataMap["distance"][i])))
	}

	product := 1
	for _, num := range cache {
		product = product * num
	}
	fmt.Println("Part 1 Product brute force: ", product)

	t := part2Cache[0]
	n := part2Cache[1]

	startBrute := time.Now()
	brute:= len(findPossibleHoldTimesBruteForce(t, n))
	elapsedBrute := time.Since(startBrute)
	
	fmt.Println("Part 2 brute force: ", brute)
	fmt.Println("Part 2 brute force execution time: ", elapsedBrute)

	startMath := time.Now()
	math := findPossibleHoldTimesMath(t, n)
	elapsedMath := time.Since(startMath)
	
	fmt.Println("Part 2 math solution: ", math)
	fmt.Println("Part 2 math solution execution time: ", elapsedMath)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
func findPossibleHoldTimesBruteForce(t int, n int) []int {

	// t: total time
	// n: current record
	// create a function L(T) = T * (t-T)
	// L: created distance
	// T: milimeters per second AND miliseconds waited

	var results []int

	for T := 1; T <= t; T++ {
		L := -T*T + t*T
		if L > n {
			results = append(results, T)
		}
	}
	return results
}

func findPossibleHoldTimesMath(t int, n int) int {

	// t: total time
	// n: current record
	// create a function L(T) = T * (t-T)
	// L: created distance
	// T: milimeters per second AND miliseconds waited

	// find intersection points of L(T) = -T^2 + t * T and L(n) = n
	// -T^2 + t * T = n
	// rearranging the equation to a quadratic form: -T^2 + t * T - n = 0
	// solving the quadratic equation to find the roots

	a := -1
	b := t
	c := -n

	delta := b*b - 4*a*c
	if delta < 0 {
		// No intersection points
		fmt.Println("No intersection points")
		return 0
	}

	sqrtDelta := math.Sqrt(float64(delta))
	T1 := (-float64(b) + sqrtDelta) / (2 * float64(a))
	T2 := (-float64(b) - sqrtDelta) / (2 * float64(a))

	// calculate the distance between the intersection points (length of the range)
	distance := int(math.Abs(T2 - T1))

	return distance
}
