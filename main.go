package main

// := combines both defalcation and initialization
// uses for local scope and not for global scopes, in global scopes mainly uses var
func main() {

	// we have two data types arrays and slices
	// arrays has fixed length and can't have different data types
	// slices can expand and shrink
	cards := newCards()
	// to add element we use append ( which will return new slice)
	cards.print()

}
