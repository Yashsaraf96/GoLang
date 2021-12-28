package main

func main() {
	filename := "my_cards"
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile(filename)
	cards := newDeckFromFile(filename)
	cards.shuffle()
	hand, cards := deal(cards, 5)
	hand.print()
	cards.print()
}
