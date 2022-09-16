package main

func main() {
	cards := newDeckFromFile("my_cards_deck")
	shuffledCards := cards.shuffle()
	shuffledCards.print()
}
