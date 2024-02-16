package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	
)

// Create a new type of deck which is a slice of string
type deck []string

func (d deck) print() {

	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Club", "Heart", "Diamond"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, suit+" "+value)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName,[]byte(d.toString()),0666)
}

func loadDeckFromFile(fileName string) deck {
     bs,err :=ioutil.ReadFile(fileName)
	 if err != nil {
		fmt.Println("Error:",err)
		os.
	 }
}