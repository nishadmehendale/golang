package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "nigel poulton"
	course := "go fundamentals"
	fmt.Println("Author:", author, "\nCourse:", course)
	fmt.Println(converter(author, course))
}

func converter(author, course string) (a, c string) {
	author = strings.ToUpper(author)
	course = strings.Title(course)
	return author, course
}
