package main

import "fmt"

func main() {
	card := deck{newCard(), "six of spades"}
	card = append(card, "queen of hearts")
	fmt.Println(card)

	card.print()

}

func newCard() string {
	return "Ace of diamonds"
}
