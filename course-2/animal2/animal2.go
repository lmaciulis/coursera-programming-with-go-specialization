package main

import (
	"errors"
	"fmt"
)

const (
	animalCow   = "cow"
	animalBird  = "bird"
	animalSnake = "snake"

	commandNew   = "newanimal"
	commandQuery = "query"

	actionEat   = "eat"
	actionMove  = "move"
	actionSpeak = "speak"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

type Bird struct {
}

type Snake struct {
}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

var animals = make(map[string]Animal)

func main() {
	var arg1, arg2, arg3 string

	for true {
		fmt.Println("Please enter command")
		fmt.Print(">")
		_, _ = fmt.Scan(&arg1)

		if arg1 != commandNew && arg1 != commandQuery {
			fmt.Printf("One of the following commands available: %s, %s. Try again.", commandNew, commandQuery)
			continue
		}

		fmt.Println("Please enter animal name")
		fmt.Print(">")
		_, _ = fmt.Scan(&arg2)

		if arg1 == commandQuery {
			animal, ok := animals[arg2]
			if !ok {
				fmt.Println("Animal with such name do not exists, please try again...")
				continue
			}

			fmt.Println("Please enter animal action")

			fmt.Print(">")
			_, _ = fmt.Scan(&arg3)

			err := callAction(arg3, animal)

			if err != nil {
				fmt.Println(err.Error())
			}

			continue
		}

		fmt.Println("Please enter animal type")
		fmt.Print(">")
		_, _ = fmt.Scan(&arg3)

		animal, err := animalFactory(arg3)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		animals[arg2] = animal
		fmt.Println("Created it!")
	}
}

func animalFactory(input string) (Animal, error) {
	switch input {
	case animalCow:
		return Cow{}, nil
	case animalBird:
		return Bird{}, nil
	case animalSnake:
		return Snake{}, nil
	default:
		return nil, errors.New("Unknown animal type, please retry...")
	}
}

func callAction(input string, animal Animal) error {
	switch input {
	case actionEat:
		animal.Eat()
		return nil
	case actionMove:
		animal.Move()
		return nil
	case actionSpeak:
		animal.Speak()
		return nil
	default:
		return errors.New("Unknown animal action, please retry...")
	}
}
