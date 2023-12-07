package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("inputs/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var hands []hand
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		n, _ := strconv.Atoi(fields[1])
		h := hand{
			cards:    fields[0],
			bid:      n,
			handType: getHandType(fields[0]),
		}
		hands = append(hands, h)
	}

	compare := func(a, b hand) int {
		if a.handType < b.handType {
			return -1
		}
		if a.handType > b.handType {
			return 1
		}
		for i := range a.cards {
			c := compareCards(rune(a.cards[i]), rune(b.cards[i]))
			if c != 0 {
				return c
			}

		}
		return 0
	}
	slices.SortFunc(hands, compare)

	sum := 0
	for pos, hand := range hands {
		multiplier := pos + 1
		sum += multiplier * hand.bid
	}
	fmt.Println(sum)

}

type hand struct {
	cards    string
	bid      int
	handType handType
}

type handType int

const (
	FiveOfAKind  handType = 7
	FourOfAKind           = 6
	FullHouse             = 5
	ThreeOfAKind          = 4
	TwoPair               = 3
	OnePair               = 2
	HighCard              = 1
)

func getHandType(cards string) handType {
	freq := map[rune]int{}
	for _, c := range cards {
		if count, ok := freq[c]; ok {
			count++
			freq[c] = count
		} else {
			freq[c] = 1
		}
	}

	for _, val := range freq {
		if val == 5 {
			return FiveOfAKind
		}
		if val == 4 {
			return FourOfAKind
		}
		if val == 3 {
			if len(freq) == 2 {
				return FullHouse
			} else {
				return ThreeOfAKind
			}
		}
	}

	count := 0
	for _, v := range freq {
		if v == 2 {
			count++
		}
	}
	if count == 2 {
		return TwoPair
	} else if count == 1 {
		return OnePair
	}

	return HighCard
}

func compareCards(a, b rune) int {
	cards := []rune{
		'A',
		'K',
		'Q',
		'J',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
	}
	if a == b {
		return 0
	}

	aIndex, bIndex := slices.Index(cards, a), slices.Index(cards, b)
	if aIndex < bIndex {
		return 1
	}
	return -1
}
