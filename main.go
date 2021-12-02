package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to advent of code!")
	dec01a, dec01b := December01()
	fmt.Printf("the answer to dec 01 is: %v, %v\n", dec01a, dec01b)
}

func December01() (int, int) {
	// open the file
	f, err := os.Open("december01.txt")
	if err != nil {
		panic("unable to open file :(")
	}
	fmt.Printf("successfully opened file!")
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

	fmt.Printf("the file is %v lines long \n", len(depths))

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