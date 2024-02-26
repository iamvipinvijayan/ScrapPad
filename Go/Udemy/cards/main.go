package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	remainingCards.print()
	fmt.Println("********************************************************")
	println(hand.toString())
	hand.shuffle().shuffle().print()
	fmt.Println("********************************************************")
}
