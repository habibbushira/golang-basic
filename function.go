package main

import (
	"fmt"
	"math"
)

func sayGreating(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Good by %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreating("Mario")
	sayGreating(("Alex"))
	sayBye(("Andrew"))

	cycleNames([]string{"Alice", "Bob", "John"}, sayGreating)
	cycleNames([]string{"Alice", "Bob", "John"}, sayBye)

	a1 := circleArea((10.5))
	a2 := circleArea((15))

	fmt.Println("Circle Area", a1, a2)
	fmt.Printf("Circle One is %0.3f and Circle Two is %0.3f\n", a1, a2)
}
