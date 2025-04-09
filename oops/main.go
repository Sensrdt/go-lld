package main

import "fmt"

// abstraction
type Animal interface {
	GetName() string
	Speak() string
}

// super class
type BaseAnimal struct {
	Name string
}

// equivalent constructor
func AssignNewAnimal(name string) BaseAnimal {
	return BaseAnimal{Name: name}
}

// getter
func (b BaseAnimal) GetName() string {
	return b.Name
}

// setter
func (b *BaseAnimal) SetName(name string) {
	b.Name = name
}

type Dog struct {
	BaseAnimal
	Breed string
}

// equivalent to @override
func (d Dog) Speak() string {
	return fmt.Sprintf("Dog %s of breed %s is barking!", d.Name, d.Breed)
}

type Cat struct {
	BaseAnimal
	Color string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("Cat %s is of colour %s", c.Name, c.Color)
}

// compile time checking if Dog is implementing methods of Animal or not
var _ Animal = (*Dog)(nil)
var _ Animal = (*Cat)(nil)

func Introduce(a Animal) {
	fmt.Println(a.Speak())
	fmt.Println(a.GetName())
}

func main() {
	dog := &Dog{
		BaseAnimal: AssignNewAnimal("Bruno"),
		Breed:      "German Shepherd",
	}

	cat := &Cat{
		BaseAnimal: AssignNewAnimal("Misty"),
		Color:      "White",
	}

	Introduce(dog)
	Introduce(cat)
}
