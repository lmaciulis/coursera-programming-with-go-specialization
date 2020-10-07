package main

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	exitChar = "X"
	sliceCap = 3
)

func main() {
	var inp string
	s := make([]int, 0, sliceCap)

	fmt.Println("Please enter an integer:")

	for true {
		_, err := fmt.Scan(&inp)
		if err != nil {
			fmt.Println("Error while reading input, please try again")
			continue
		}

		if inp == exitChar {
			break
		}

		i, err := strconv.Atoi(inp)
		if err != nil {
			fmt.Println("Error while converting input to integer, please try again")
			continue
		}

		s = append(s, i)
		sort.Ints(s)

		printSlice(s)

		fmt.Println("Please enter next integer or type 'X' to exit")
	}
}

func printSlice(s []int) {
	fmt.Printf("%v\n", s)
}
