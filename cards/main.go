package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// Declaring a function named newCard () that will return "string"
func newCard() string {
	return "Ace of Spades"
}
