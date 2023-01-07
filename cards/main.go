package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("----------")
	// remainingCards.print()
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}
