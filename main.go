package main

import "fmt"

func main() {
	// `var card string = "Ace of Spades"`` is equal to `card := "Ace of Spades"``
	card := "Ace of Spades"   // := is only use to define/intialize new variable
	card = "Five of Diamonds" // to update just use = only

	fmt.Println(card)
}
