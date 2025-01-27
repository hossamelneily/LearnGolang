package main

import (
	"fmt"
)

// := combines both defalcation and initialization
// uses for local scope and not for global scopes, in global scopes mainly uses var
func main() {
	card := newCard()
	fmt.Printf(card)

}

// need to specify the type of return value
func newCard() string {
	return "Five of diamonds"
}
