package main

import (
	"fmt"
)

// Variables declared outside the function
// should be declared with variables block
// Variable names cannot be any of the go keywords
// https://go101.org/article/keywords-and-identifiers.html#:~:text=Keywords%201%20const%2C%20func%2C%20import%2C%20package%2C%20type%20and,used%20to%20control%20flow%20of%20code.%20More%20items
// No spaces and special characters
var (
	name         string
	course       string
	module, clip int
)

func main() {
	fmt.Println("Name and Course Set to ", name, "and", course, ".")
	fmt.Println("Module and Clip Set to ", module, "and", clip, ".")
}
