package main

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
	for _, card := range d {
		println(card)
	}
}
