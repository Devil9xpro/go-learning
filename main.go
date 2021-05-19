package main

func main() {
	//cards := newDeck()
	//hand, remainingCard := deal(cards, 5)
	//hand.print()
	//remainingCard.print()

	//cards := newDeckFromFile("my_card")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}