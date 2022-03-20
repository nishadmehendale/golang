package main

import "fmt"

func main() {
	const c = "Hello"
	fmt.Println(c, "Pluralsight!")

	/*
		Uncomment the following code to get an error
		Since constants cannnot be changed
	*/
	// c = "Namaskar"
	//fmt.Println(c, "Pluralsight!")
}
