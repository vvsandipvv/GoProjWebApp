package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	faces := []string{"Clubs", "Spades", "Diamonds", "Hearts"}
	values := []string{"Ace", "One", "Two", "Three"}

	for _, k := range faces {
		for _, l := range values {
			cards = append(cards, k+" of "+l)
		}
	}

	return cards
}

func (d deck) print() {
	for i, j := range d {
		fmt.Println(i, j)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
