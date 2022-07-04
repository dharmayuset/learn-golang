package main

import "fmt"

func main() {
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Diamonds")
	// its example slice and for in go language
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
func newCard() string {
	return "Five of Diamonds"
}
