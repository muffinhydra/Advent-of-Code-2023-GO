package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"
)

type Transfer struct {
	inputStart  int
	inputEnd    int
	outputStart int
	outputEnd   int
	r           int
	seedMap     string //mostly for debug purposes
}

func main() {

	// Open the txt file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var seedArray []int
	var mapSpec string
	var mapSpecArray []string

	seedMaps := make(map[string][]Transfer)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		fields := strings.Fields(scanner.Text())

		switch {

		case len(fields) == 0:
			continue

		case fields[0] == "seeds:":
			for i := 1; i < len(fields); i++ {
				num, err := strconv.Atoi(fields[i])
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				seedArray = append(seedArray, num)
			}
			continue

		case !unicode.IsDigit([]rune(fields[0])[0]):
			mapSpec = fields[0]
			mapSpecArray = append(mapSpecArray, fields[0])
			fmt.Println("created new Map", fields[0])
			continue

		case unicode.IsDigit([]rune(fields[0])[0]):
			r, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			input, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			output, err := strconv.Atoi(fields[0])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			var tempKey Transfer
			tempKey.inputStart = input
			tempKey.inputEnd = input + r
			tempKey.outputStart = output
			tempKey.outputEnd = output + r
			tempKey.r = r
			tempKey.seedMap = mapSpec
			seedMaps[mapSpec] = append(seedMaps[mapSpec], tempKey)

		}
		fmt.Println("finished Line", fields)
	}

	fmt.Println("finished building maps")

	fmt.Println("traversing maps")

	totalSeeds := 0
	for i := 1; i < len(seedArray); i += 2 {
		totalSeeds += seedArray[i]
	}

	seedCounter := int64(0) 
	var progressMutex sync.Mutex

	seedCatcher := math.MaxInt64
	var mutex sync.Mutex
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < len(seedArray); i += 2 {
		wg.Add(1)
		go func(start, count int) {
			defer wg.Done()
			processSeeds(start, count, mapSpecArray, seedMaps, &seedCatcher, &mutex, &seedCounter, &progressMutex)
		}(seedArray[i], seedArray[i+1])
	}

	// Display progress with timestamps 
	go func() {
		for {
			progressMutex.Lock()
			processed := atomic.LoadInt64(&seedCounter)
			progressMutex.Unlock()
			elapsed := time.Since(startTime)
			fmt.Printf("[%s] Processed %d out of %d seeds\n", elapsed.Round(time.Second), processed, totalSeeds)
			time.Sleep(time.Second) 
		}
	}()


	wg.Wait()

	
	progressMutex.Lock()
	defer progressMutex.Unlock()
	elapsed := time.Since(startTime)
	fmt.Printf("[%s] Processing complete. Processed %d out of %d seeds\n", elapsed.Round(time.Second), atomic.LoadInt64(&seedCounter), totalSeeds)
	fmt.Println("Minimum:", seedCatcher)
}

func processSeeds(start, count int, mapSpecArray []string, seedMaps map[string][]Transfer, seedCatcher *int, mutex *sync.Mutex, seedCounter *int64, progressMutex *sync.Mutex) {
	localSeedCatcher := math.MaxInt64
	for j := 0; j < count; j++ {
		seed := start + j
		for _, seedSpec := range mapSpecArray {
			for _, seedTransfer := range seedMaps[seedSpec] {
				if seed >= seedTransfer.inputStart && seed <= seedTransfer.inputEnd {
					delta := seed - seedTransfer.inputStart
					seed = seedTransfer.outputStart + delta
					break
				}
			}
		}
		if seed < localSeedCatcher {
			localSeedCatcher = seed
		}

	
		atomic.AddInt64(seedCounter, 1)
	}

	mutex.Lock()
	defer mutex.Unlock()
	if localSeedCatcher < *seedCatcher {
		*seedCatcher = localSeedCatcher
	}
}
