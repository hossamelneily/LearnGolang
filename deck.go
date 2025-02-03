package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// go is not a oo programming language
// Create a new type of 'deck'
// which is a slice of strings

// that means a new class that extends the class string
type deck []string

func newCards() deck {
	cards := deck{}
	cardSuites := []string{"Ace", "Diamond"}
	cardValues := []string{"One", "Two", "Three"}

	for _, suite := range cardSuites {
		for _, card := range cardValues {
			cards = append(cards, suite+" of "+card)
		}
	}
	return cards
}

// receiver means any variable of type deck now gets access to the print method
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) saveToFile(filename string) error {
	toString := strings.Join([]string(d), ",")
	return os.WriteFile(filename, []byte(toString), 0666)
}

func readFromFile(filename string) []string {

	byteFromFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	toString := deck(strings.Split(string(byteFromFile), ","))
	return toString
}

func (d deck) suffle() deck {
	for i, card := range d {
		randomNumber := rand.Intn(len(d) - 1)
		d[i] = d[randomNumber]
		d[randomNumber] = card

	}
	return d
}
