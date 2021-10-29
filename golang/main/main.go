package main

import "fmt"

type Person struct {
	name string
	edad int
}

func (p Person) greet() {
	fmt.Print("Hello, " + p.name)
}

func main() {

	p1 := Person{"Dylan", 18}

	p1.greet()

}
