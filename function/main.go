package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Diamonds")

	cards.print()
}
func newCard() string {
	return "Five of Diamonds"
}
