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
		s := strings.Fields(line)
		n, _ := strconv.Atoi(s[1])
		h := hand{
			cards:    s[0],
			bid:      n,
			handType: getHandTypeWithJokers(s[0]),
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
	for i, h := range hands {
		multiplier := i + 1
		sum += multiplier * h.bid
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
	for _, card := range cards {
		if count, ok := freq[card]; ok {
			count++
			freq[card] = count
		} else {
			freq[card] = 1
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

func getHandTypeWithJokers(cards string) handType {
	freq := map[rune]int{}
	for _, card := range cards {
		if count, ok := freq[card]; ok {
			count++
			freq[card] += 1
		} else {
			freq[card] = 1
		}
	}

	jokers := freq['J']
	if jokers == 0 {
		return getHandType(cards)
	}

	var mostCommonCard rune
	var maxVal int
	for card, val := range freq {
		if card == 'J' {
			continue
		}
		if val > maxVal {
			mostCommonCard = card
			maxVal = val
		}
	}

	cards = strings.ReplaceAll(cards, "J", string(mostCommonCard))
	return getHandType(cards)
}

func compareCards(a, b rune) int {
	cards := []rune{
		'A',
		'K',
		'Q',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
		'J',
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
