package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	for i := 0; i < 2; i++ {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error while reading input")
			continue
		}

		if found(sanitize(input)) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}

func found(input string) bool {
	return strings.Contains(input, "a") &&
		strings.HasPrefix(input, "i") &&
		strings.HasSuffix(input, "n")
}

func sanitize(input string) string {
	s := strings.TrimSpace(input)
	s = strings.ToLower(s)

	return s
}
