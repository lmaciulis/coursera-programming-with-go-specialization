package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	for i := 0; i < 2; i++ {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error while reading input")
			continue
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error while parsing float from input")
		}

		fmt.Println(int(num))
	}
}
