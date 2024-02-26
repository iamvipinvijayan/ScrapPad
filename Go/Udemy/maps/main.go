package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var color map[string]string
	// fmt.Println(color)
	// col := make(map[string]string)
	// fmt.Println(col)
	// // fmt.Println(colors)
	// fmt.Println(colors["red"])
	printMap(colors)

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("The Key for color :", color, " is ", hex)
	}

}
