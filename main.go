package main

import "fmt"

func newFunc(random int) int {
	return random * 2
}
func main() {
	fmt.Println("Hello World!")
	fmt.Println(newFunc(2))
}
