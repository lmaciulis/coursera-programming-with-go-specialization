package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

const (
	maxInputLen = 20
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Starting a program...")
	fmt.Println("Please enter a file name:")

	var inp string
	var data []Name

	_, err := fmt.Scan(&inp)
	if err != nil {
		log.Fatal(errors.New("error while reading input, please try again"))
	}

	filename, err := getFilePath(inp)
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		name, err := getName(text)
		if err != nil {
			fmt.Println("Error while converting line to struct. Skipping line...")
			fmt.Println(err.Error() + ", Input line: " + text)
			continue
		}

		data = append(data, name)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Outputting collected data:\n")

	for _, v := range data {
		fmt.Printf("%s %s\n", v.fname, v.lname)
	}

	fmt.Println("\nProgram finished execute it's code")
}

func getFilePath(input string) (string, error) {
	if filepath.IsAbs(input) {
		return input, nil
	}

	absPath, err := filepath.Abs(input)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

func getName(input string) (Name, error) {
	var name Name

	words := strings.Fields(input)
	if len(words) != 2 {
		return name, errors.New("Line should contain 2 words")
	}

	name.fname = sanitizeInput(words[0])
	name.lname = sanitizeInput(words[1])

	return name, nil
}

func sanitizeInput(input string) string {
	if utf8.RuneCountInString(input) > maxInputLen {
		rs := []rune(input)

		return string(rs[:maxInputLen])
	}

	return input
}
