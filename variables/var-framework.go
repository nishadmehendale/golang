package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Variables declared outside the function
// should be declared with variables block
// Variable names cannot be any of the go keywords
// https://go101.org/article/keywords-and-identifiers.html#:~:text=Keywords%201%20const%2C%20func%2C%20import%2C%20package%2C%20type%20and,used%20to%20control%20flow%20of%20code.%20More%20items
// No spaces and special characters
var (
	name   string
	course string
	module = "4"
	clip   = 3
)

func main() {
	// Local Variables scope is only limited to the function it is declared within
	courseComplete := false
	feedback := "Awesome"
	fmt.Println("Name and Course Set to ", name, "and", course, ".")
	fmt.Println("Module and Clip Set to ", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
	// Conversion of sting to int
	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Total of Module and Clip", total)
	}
	fmt.Println("Is Course Complete?", courseComplete)
	// Printing Memeory address using '&'
	fmt.Println("Memory Address of Course Variable", &course)

	// When creating a variable with '*' makes it a pointer variable
	// When accessing the pointer variable with '*' dereferences it and returns actual value
	var ptr *string = &module
	fmt.Println("Pointing Module Variable At Address,", ptr, "which holds the value,", *ptr)

	fmt.Println("Feedback is", feedback)
	// this is a pass by value function it doesnt update the original clip variables value
	updateFeedback(feedback)
	fmt.Println("Updated Feedback post pass by Value", feedback)
	updateFeedbackReally(&feedback)
	fmt.Println("Updated Feedback post pass by Value", feedback)

}

func updateFeedback(feedback string) string {
	feedback = "Amazing"
	fmt.Println("Updating feedback to ", feedback, " but not really")
	return feedback
}

func updateFeedbackReally(feedback *string) string {
	*feedback = "Amazing"
	fmt.Println("Updating feedback to ", *feedback, " really")
	return *feedback
}
