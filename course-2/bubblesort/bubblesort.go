package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const maxInputArrLen = 10

func main() {
	fmt.Println("This program sorts array of integers, using BubbleSort")

	for true {
		fmt.Println("Please enter space separated integers:")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		}

		arr := parseIntArr(input)
		BubbleSort(&arr)

		fmt.Println("Outputting sorted integers array:")
		fmt.Printf("\n%v\n\n", arr)
	}
}

func BubbleSort(arr *[]int) {
	n := len(*arr)

	for i := 0; i < n; i++ {
		// optimization for not repeating loops when array is already sorted
		swapCalled := false

		for j := 0; j < n-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				Swap(arr, j)
				swapCalled = true
			}
		}

		if !swapCalled {
			break
		}
	}
}

func Swap(arr *[]int, index int) {
	(*arr)[index], (*arr)[index+1] = (*arr)[index+1], (*arr)[index]
}

func parseIntArr(input string) []int {
	out := make([]int, 0, maxInputArrLen)

	for _, v := range strings.Fields(input) {
		if len(out) == maxInputArrLen {
			fmt.Printf("Max input length is %d. Skipping leading inputs...\n", maxInputArrLen)

			return out
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Error while converting characters '%s' to integer, skipping...\n", v)
			continue
		}

		out = append(out, i)
	}

	return out
}
