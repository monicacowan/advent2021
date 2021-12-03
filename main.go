package main

import (
	"fmt"
	"github.com/monicacowan/advent2021/dec01"
	"github.com/monicacowan/advent2021/dec02"
)

func main() {
	fmt.Println("Welcome to advent of code!")

	dec01a, dec01b := dec01.December01()
	fmt.Printf("the answer to dec 01 is: %v, %v\n", dec01a, dec01b)

	dec02a, dec02b := dec02.December02()
	fmt.Printf("the answer to dec 02 is: %v, %v\n", dec02a, dec02b)
}