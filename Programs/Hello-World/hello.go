package main

import "fmt"

// simple hello world program
func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Hello " + name + "!")
}
