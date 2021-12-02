package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to advent of code!")
	dec01a, dec01b := December01()
	fmt.Printf("the answer to dec 01 is: %v, %v\n", dec01a, dec01b)
	dec02a := December02()
	fmt.Printf("the answer to dec 02 is: %v\n", dec02a)
}

func December01() (int, int) {
	// open the file
	f, err := os.Open("december01.txt")
	if err != nil {
		panic("unable to open file :(")
	}
	 // read into an array
	 var depths []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("couldn't convert string to int")
		}
		depths = append(depths, num)
	}

	countA := 0
	// part a
	for index, item := range depths {
		if index == 0 {
			continue
		}

		// check to see if the current item is bigger than the previous one
		if item > depths[index-1] {
			countA++
		}
	}

	// part b
	countB := 0
	for i, depth := range depths {
		if i > len(depths) - 4 {
			continue
		}

		setA := depth + depths[i+1] + depths[i+2]
		setB := depths[i+1] + depths[i+2] + depths[i+3]
		if setB > setA {
			countB++
		}
	}

	err = f.Close()
	if err != nil {
		panic("unable to close file :(")
	}
	return countA, countB
}

func December02() int {
	// open the file
	file, err := os.Open("december02.txt")
	if err != nil {
		panic("couldn't open the file :(\n")
	}

	type movement struct{
		direction string
		distance int
	}
	var movements []movement
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		 m := strings.Fields(scanner.Text()) // will be in the format: direction distance

		 // convert distance to int32
		 dist, err := strconv.Atoi(m[1])
		 if err != nil {
			 panic("couldn't convert string to int")
		 }

		 movements = append(movements, movement{
			 direction: m[0],
			 distance:  dist,
		 })
	}

	// initialize all three categories
	forward := 0
	up := 0
	down := 0

	for _, mvmt := range movements {
		switch mvmt.direction {
		case "forward":
			forward += mvmt.distance
		case "up":
			up += mvmt.distance
		case "down":
			down += mvmt.distance
		}
	}

	depth := down - up

	return depth * forward
}