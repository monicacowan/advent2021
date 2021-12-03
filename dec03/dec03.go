package dec03

import (
	"bufio"
	"math"
	"os"
)

func Decmeber03() int {
	// open file and read in the contents
	f, err := os.Open("input/december03.txt")
	if err != nil {
		panic("couldn't open the file :(\n")
	}

	// read into an array
	var diagnostics []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		diagnostic := scanner.Text()
		diagnostics = append(diagnostics, diagnostic)
	}

	// determine the length of the diagnostic item
	diagLength := len(diagnostics[0])

	// initialize gamma and epsilon
	var gamma []int
	var epsilon []int

	// for each position in the diagnostic value (ie, 0 -> diagLength)...
	// make a count of 0's and 1's
	// max gets appended to gamma, other gets appended to epsilon
	zeroes := 0
	ones := 0
	for i := 0; i < diagLength; i++ {
		zeroes = 0
		ones = 0
		for _, d := range diagnostics {
			if d[i] == 48 { // we've got ASCII on our hands
				zeroes++
			} else {
				ones++
			}
		}
		if ones > zeroes {
			gamma = append(gamma, 1)
			epsilon = append(epsilon, 0)
		} else {
			gamma = append(gamma, 0)
			epsilon = append(epsilon, 1)
		}
	}

	return binaryCalculator(gamma) * binaryCalculator(epsilon)
}

// takes in a binary number represented as a slice of ints, returns the binary number as an int
func binaryCalculator(input []int) int {
	num := 0
	exp := 0.0
	binaryLength := len(input)

	// start at the end of the array and head back to the start
	for i := binaryLength - 1; i > -1; i-- {
		num += input[i] * int(math.Pow(2, exp))
		exp++
	}
	return num
}
