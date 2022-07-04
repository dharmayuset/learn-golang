package main

import "fmt"

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Diamonds")
	// its example slice and for in go language
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
func newCard() string {
	return "Five of Diamonds"
}
