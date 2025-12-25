package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println("Animal Speak")
}

type Dog struct {
	Animal
	Breed string
}

func (d *Dog) Speak() {
	fmt.Println("Dog Speak")
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Labrador",
	}
	fmt.Println(d.Name)
	fmt.Println(d.Breed)
	d.Speak()
	d.Animal.Speak()
}
