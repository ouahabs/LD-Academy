package main

import (
	"fmt"
	"os"
	"strconv"
)

// This function checks if we have at least two arguments
// and then converts the first two to integers
// otherwise fails with an error code
func sanity() int {
	args := os.Args
	if len(args) <= 2 {
		return 0
	}
	if _, err := strconv.Atoi(args[1]); err != nil {
		return 1
	}
	if _, err := strconv.Atoi(args[2]); err != nil {
		return 2
	}
	return 10
}

// This function gets the first two arguments
// and returns them as integers
func getArgs() (int, int) {
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])

	return x, y
}

// This function uses euclidean algorithm to divide a % b
// until the remainder is 0, in which case we get GCF(b, 0)
// which is equal to b
func GCF(a, b int) int {
	if a == 0 && b == 0 {
		return -1
	}

	// Find recuresively the GCF
	if b != 0 {
		return GCF(b, a%b)
	} else {
		return a
	}

}

func main() {
	test := sanity()
	switch test {
	case 0:
		fmt.Println("Not enough arguments, please include at least two")
	case 1:
		fmt.Printf("%s isn't an integer \n", os.Args[1])
		os.Exit(1)
	case 2:
		fmt.Printf("%s isn't an integer \n", os.Args[2])
		os.Exit(1)
	case 10:
		a, b := getArgs()
		fmt.Printf("Greatest common factor (GCF) between %d and %d is %d \n", a, b, GCF(a, b))
	default:
		fmt.Println("Failed in getting arguments")
	}
}
