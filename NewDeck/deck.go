package main
import "fmt"
// create new deck type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clover"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, suit+" of "+values)
		}
	}
	return cards
}

//function receiver from main.go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i,card)
	}
}
func deal(d deck, handSize int) (deck,deck) {
	return d[:handSize], d[handSize:]
}
