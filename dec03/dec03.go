package dec03

import (
	"bufio"
	"math"
	"os"
)

func Decmeber03() (int, int) {
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
	for i := 0; i < diagLength; i++ {
		ones, zeroes := bitCount(diagnostics, i)
		if ones > zeroes {
			gamma = append(gamma, 1)
			epsilon = append(epsilon, 0)
		} else {
			gamma = append(gamma, 0)
			epsilon = append(epsilon, 1)
		}
	}

	// part two: search for the oxygen generator rating
	// want the number with the most common bit per position!

	commonBitASCII, leastCommonBitASCII := 48, 48

	// initialize the list to diagnostics; this will change as we return a filtered list in the loop
	list := diagnostics

	//testList := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	for i := 0; i < diagLength; i++ {
		// get the common bit
		commonBitASCII = 48
		ones, zeroes := bitCount(list, i)
		if zeroes < ones || zeroes == ones {
			commonBitASCII = 49
		}
		list = filter(list, i, commonBitASCII, bitPresent)
		if len(list) == 1 {
			break
		}
	}

	oxygenGeneratorRating := list[0]

	// now the co2 scrubber rating: want the bit with the least common bit present

	list2 := diagnostics

	for i := 0; i < diagLength; i++ {
		// get the least common bit
		// assume that zero is the least common bit
		leastCommonBitASCII = 48
		ones, zeroes := bitCount(list2, i)

		if zeroes > ones {
			leastCommonBitASCII = 49
		}
		list2 = filter(list2, i, leastCommonBitASCII, bitPresent)
		if len(list2) == 1 {
			break
		}
	}
	co2scrubberRating := list2[0]

	oxyArray := stringOfBinaryNumbersToArray(oxygenGeneratorRating)
	co2Array := stringOfBinaryNumbersToArray(co2scrubberRating)

	return binaryCalculator(gamma) * binaryCalculator(epsilon), binaryCalculator(oxyArray) * binaryCalculator(co2Array)
}

func stringOfBinaryNumbersToArray(s string) []int {
	var arr []int
	for _, n := range s {
		if n == 48 {
			arr = append(arr, 0)
		} else {
			arr = append(arr, 1)
		}
	}
	return arr
}

func filter(diagnostics []string, i int, commonBit int, test func(string, int, int) bool) (ret []string) {
	backup := diagnostics
	for _, d := range diagnostics {
		if test(d, i, commonBit) {
			ret = append(ret, d)
		}
	}

	if len(ret) == 0 {
		ret = backup
	}
	return ret
}

// this test takes in a binary number represented as a string, and a digit location.  If the number has the bit in that digit
// location, returns true.
func bitPresent(d string, i int, commonBit int) bool {
	return d[i] == uint8(commonBit)
}

// takes in an array of binary numbers represented as strings, and a digit location, returns the number of instances of ones and zeroes
func bitCount(diagnostics []string, i int) (int, int) {
	zeroes := 0
	ones := 0
	for _, d := range diagnostics {
		if d[i] == 48 { // we've got ASCII on our hands
			zeroes += 1
		} else {
			ones += 1
		}
	}
	return ones, zeroes
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
