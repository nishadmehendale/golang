package main

import (
	"fmt"
	"os"
)

func main() {
	// Lists Environment Variables in KV pair
	fmt.Println(os.Environ())

	// Lists Environment Variables in KV pair but each one on the new line
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// use USER in case of Mac or Linux
	user := os.Getenv("USERNAME")
	fmt.Println(user, "is logged in")
}
