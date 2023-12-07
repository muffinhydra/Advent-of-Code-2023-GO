package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bet   int
	rank  int
}

type ByRank []Hand

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {

	m := a[i]
	n := a[j]

	numbersMap := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 0,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}
	switch {
	case m.rank < n.rank:
		return true
	case m.rank > n.rank:
		return false
	case m.rank == n.rank:
		for i := 0; i < len(m.cards); i++ {

			if numbersMap[string(m.cards[i])] == numbersMap[string(n.cards[i])] {
				continue
			}
			if numbersMap[string(m.cards[i])] < numbersMap[string(n.cards[i])] {
				return true
			} else {
				return false
			}
		}
	}
	return false // only so that the compiler can shut up about it 
}

func evaluateRank(hand Hand) int {
	cards := strings.Split(hand.cards, "")
	sort.Strings(cards)

	counts := make(map[string]int)
	for _, card := range cards {
		counts[card]++
	}

	wildcardCount := counts["J"]
	delete(counts, "J") // remove jokers from count for evaluating hand
	if len(counts) == 0 {
		return 7
	}

	frequencies := make([]int, 0, len(counts))
	for _, count := range counts {
		frequencies = append(frequencies, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequencies))) //sort to have the 2 highest frequencies at frequencies[0] and frequencies[1]

	switch {
	case frequencies[0]+wildcardCount == 5:
		return 7 // Five of a kind
	case frequencies[0]+wildcardCount == 4:
		return 6 // Four of a kind
	case frequencies[0]+wildcardCount == 3 && frequencies[1] == 2:
		return 5 // Full house
	case frequencies[0]+wildcardCount == 3 && frequencies[1] == 1:
		return 4 // Three of a kind
	case frequencies[0]+wildcardCount == 2 && frequencies[1] == 2:
		return 3 // Two pair
	case frequencies[0]+wildcardCount == 2 && frequencies[1] == 1:
		return 2 // One pair
	default:
		return 1 // High card
	}
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

	var sum int
	var hands []Hand

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Fields(line)
		var hand Hand
		hand.cards = splits[0]
		hand.bet, _ = strconv.Atoi(splits[1])
		hand.rank = evaluateRank(hand)
		hands = append(hands, hand)
	}
	sort.Sort(ByRank(hands))

	fmt.Println("hands :", hands)
	for i := 0; i < len(hands); i++ {
		sum += hands[i].bet * (i + 1)
	}

	fmt.Println("Sum winnings :", sum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
