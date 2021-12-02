package dec01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func December01() (int, int) {
	// open the file
	f, err := os.Open("input/december01.txt")
	if err != nil {
		fmt.Print(err.Error())
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