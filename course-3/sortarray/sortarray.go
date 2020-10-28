package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter space separated integers:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}

	arr := parseIntArr(input)

	arrLen := len(arr)
	chunkLen := arrLen / 4

	if arrLen < 4 {
		fmt.Println("Please provide at least 4 integers...")
		fmt.Println("Program simply sort this array using bubble sort:")

		fmt.Println(bubbleSort(arr))

		os.Exit(0)
	}

	fmt.Println(arr)

	c1 := make(chan []int, 4)

	for i := 0; i < 4; i++ {
		if i == 3 {
			go sortRoutine(arr[i*chunkLen:], c1)
		} else {
			go sortRoutine(arr[i*chunkLen:(i+1)*chunkLen], c1)
		}
	}

	res := merge(merge(<-c1, <-c1), merge(<-c1, <-c1))

	fmt.Println(res)
}

func sortRoutine(inp []int, c chan []int) {
	fmt.Println(inp)
	c <- bubbleSort(inp)
}

func bubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}

func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res[i] = left[0]
			left = left[1:]
		} else {
			res[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		res[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		res[i] = right[j]
		i++
	}

	return res
}

func parseIntArr(input string) []int {
	out := make([]int, 0)

	for _, v := range strings.Fields(input) {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Error while converting characters '%s' to integer, skipping...\n", v)
			continue
		}

		out = append(out, i)
	}

	return out
}
