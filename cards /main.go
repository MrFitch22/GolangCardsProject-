package main

func main() {
	// making multiple cards at a time using slice

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
