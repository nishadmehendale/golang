package main

import (
	"fmt"
)

func main() {
	heightOfA := 180
	heightOfB := 160
	// Simple if else hierarchy
	if heightOfA > heightOfB {
		fmt.Println("A is taller")
	} else if heightOfB > heightOfA {
		fmt.Println("B is taller")
	} else {
		fmt.Printf("Same height")
	}
	// Value declaration at the if function
	// scope of variable limited to if hierarchy
	if heightOfC, heightOfD := 110, 180; heightOfC > heightOfD {
		fmt.Println("C is taller")
	} else if heightOfD > heightOfC {
		fmt.Println("D is taller")
	} else {
		fmt.Printf("Same height")
	}
	// Switch case
	switch "Kubernetes" {
	case "Istio":
		fmt.Println("MatchFound for Istio")
	case "Kubernetes":
		fmt.Println("MatchFound for Kubernetes")
	default:
		fmt.Println("Match Found")
	}

}
