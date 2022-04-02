package main

func main() {
	cards := newDeck()
	cards.saveToFile("deck")
	cards.shuffle()
	cards.print()
}
