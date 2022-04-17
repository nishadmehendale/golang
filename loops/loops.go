package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Simple For Loop
		timer := 10 (Pre Statement) [Optional]
		timer >= 0 (Condtion)
		timer-- (Post Statement) [Optional]
	*/
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	/*
		Looping over a collection / range
	*/
	courses := []string{"Course One", "Course Two"}
	// Prints index value
	for i := range courses {
		fmt.Println(i)
	}
	// Prints course name
	for _, i := range courses {
		fmt.Println(i)
	}
	coursesCompleted := []string{"Course One"}
	// Nested Loops
	for _, i := range courses {
		for _, j := range coursesCompleted {
			if i == j {
				break
			}
			fmt.Println("Courses In Progress", i)
		}
	}
}
