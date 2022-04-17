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
}
