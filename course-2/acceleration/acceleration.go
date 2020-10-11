package main

import "fmt"

func main() {
	fmt.Println("Starting the program...")

	var a, v, s, t float64

	fmt.Print("Enter acceleration: ")
	_, _ = fmt.Scanln(&a)

	fmt.Print("Enter initial velocity: ")
	_, _ = fmt.Scanln(&v)

	fmt.Print("Enter initial displacement: ")
	_, _ = fmt.Scanln(&s)

	fmt.Print("Enter time:")
	_, _ = fmt.Scanln(&t)

	fn := GenDisplaceFn(a, v, s)

	fmt.Printf("Displacement after %v seconds: %v", t, fn(t))
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v*t + s
	}
}
