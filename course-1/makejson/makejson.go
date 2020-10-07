package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("This program prompts you to enter person's name and address, and prints out it's JSON")
	i := 0
	for true {
		if i == 0 {
			fmt.Println("Starting a program...")
		} else {
			fmt.Println("Program restarted...")
		}

		err := loopHandler()
		if err != nil {
			fmt.Println(err.Error())
		}
		i++
	}
}

func loopHandler() error {
	reader := bufio.NewReader(os.Stdin)
	p := make(map[string]string)

	fmt.Println("Please enter 'name'")
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	p["name"] = strings.TrimSuffix(name, "\n")

	fmt.Println("Please enter 'address'")
	address, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	p["address"] = strings.TrimSuffix(address, "\n")

	j, err := json.Marshal(p)
	if err != nil {
		return err
	}

	fmt.Println("JSON object output:")
	fmt.Printf("\n%s\n\n", j)

	return nil
}
