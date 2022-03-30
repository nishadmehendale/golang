package main

import (
	"fmt"
	"math/rand"
	"time"
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
		// Uncomment below line to execute next individual case
		// fallthrough
	default:
		fmt.Println("Match Not Found")
	}

	// Idioswitch Example
	switch tmpNum := random(); tmpNum {
	case 1, 3, 5, 7, 9:
		fmt.Println("We got an odd number", tmpNum)
	case 2, 4, 6, 8, 0:
		fmt.Println("We got an even number", tmpNum)
	default:
		fmt.Println("We did not get a number", tmpNum)
	}

}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
