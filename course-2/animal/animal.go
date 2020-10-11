package main

import (
	"errors"
	"fmt"
)

const (
	animalCow   = "cow"
	animalBird  = "bird"
	animalSnake = "snake"

	actionEat   = "eat"
	actionMove  = "move"
	actionSpeak = "speak"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var an, ac string

	for true {
		fmt.Println("Please type an animal")
		fmt.Print(">")
		_, _ = fmt.Scan(&an)

		animal, err := animalFactory(an)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("Please type an action")
		fmt.Print(">")
		_, _ = fmt.Scan(&ac)

		err = callAction(ac, &animal)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func callAction(input string, animal *Animal) error {
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

func animalFactory(input string) (Animal, error) {
	switch input {
	case animalCow:
		return Animal{"grass", "walk", "moo"}, nil
	case animalBird:
		return Animal{"worms", "fly", "peep"}, nil
	case animalSnake:
		return Animal{"mice", "slither", "hsss"}, nil
	default:
		return Animal{}, errors.New("Unknown animal, please retry...")
	}
}
