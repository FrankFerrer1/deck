package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("myCards")
}

//Go has 2 basic data structure, array and slice
//Array is fixed length, a slice is an array that can grow or shrink
//array or slice must be defined by a data type, like in java, all must be the same type
