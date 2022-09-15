package main

import "fmt"

func main() {
	// card := "Ace of Spades"
	cards := []string{newCard(), "Ace of Diamonds"}
	// append do not return modified slice but return new copy of slice
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
