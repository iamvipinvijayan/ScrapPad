package main

import "fmt"

func main() {
	fmt.Println("OddOrEvenFunction")
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range numbers {
		if value%2 != 0 {
			fmt.Printf("The number %d  given is Odd\n", value)
		}
		if value%2 == 0 {
			fmt.Printf("The number %d is Even \n", value)
		}

	}
}
