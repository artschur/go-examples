package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Speak() {
	fmt.Printf("meu nome é %v", a.Name)
}

func main() {
	penguin := Animal{
		Name: "Penguin",
		Age:  18,
	}

	penguin.Speak()

}
