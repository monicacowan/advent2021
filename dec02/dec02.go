package dec02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func December02() int {
	// open the file
	file, err := os.Open("input/december02.txt")
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
