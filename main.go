package main

import "fmt"

// Returns the product of all given factors.
// In case no factor at all is passed zero is returned.
func product(factors ...int) int {
	if len(factors) == 0 {
		return 0
	}

	result := 1

	for _, factor := range factors {
		// x *= y is a short form for x = x * y
		result *= factor
	}

	return result
}

// Returns a new slice where the first occurence of target
// has been removed if present. Otherwise a slice with the
// same elements is returned.
func remove(slice []int, target int) []int {
	for index, value := range slice {
		if value == target {
			// A simple way of removing elements is reslicing.
			// It means that we cut a slice from the beginning
			// and from the cutting point and glue them together.
			return append(slice[:index], slice[index+1:]...)
		}
	}

	return slice
}

func main() {
	fmt.Println("This project contains code for testing purposes only. Please use go test instead of go run.")
}
