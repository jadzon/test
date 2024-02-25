package main

import "fmt"

type book struct {
	title string
	price float64
}

func newFunc(random int) int {
	return random * 2
}
func main() {
	fmt.Println("Hello World!")
	fmt.Println(newFunc(2))
	myBook := book{title: "Ksiazka trapu", price: 11.77}
	fmt.Printf("Cena: %.2f, tytuł: %s", myBook.price, myBook.title)
}
