package main

import "fmt"

func main() {
	//var card string = "ace of spade"
	//tes
	card := "ace of spades"
	//new assign
	card = "ace of diamonds"
	newc := newCard()
	fmt.Println(card)
	fmt.Println(newc)
}
func newCard() string {
	return "Five of Diamonds"
}
