package main

import "fmt"

// create new deck type of 'deck'
// which is a slice of strings
type deck []string

//function receiver from main.go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
