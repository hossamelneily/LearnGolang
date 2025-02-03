package main

import "fmt"

// := combines both defalcation and initialization
// uses for local scope and not for global scopes, in global scopes mainly uses var
func main() {

	// we have two data types arrays and slices
	// arrays has fixed length and can't have different data types
	// slices can expand and shrink
	cards := newCards()
	// to add element we use append ( which will return new slice)
	cards.print()

	// convert type deck to slice of strings and then to string with comma seperated
	// cards.saveToFile("test")
	// for _, card := range readFromFile("test") {
	// 	fmt.Println(card)
	// }

	cards.suffle()
	fmt.Println(cards)

}
