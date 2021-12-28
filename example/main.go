package main

import (
	"fmt"

	"github.com/langbox/handy"
)

func main() {

	id, err := handy.GenShortID()
	fmt.Println(id, err)

	str := " 1111    "
	s := handy.Trim(str)
	fmt.Println(len(s))
}
