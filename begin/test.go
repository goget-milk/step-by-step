package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Собака %s лает", d.Name)
}

func (d Dog) Bark() string {
	return fmt.Sprintf("%s громко лает!", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("Кот %s мяукает", c.Name)
}

func (c Cat) Purr() string {
	return fmt.Sprintf("%s мурлычет.", c.Name)
}

func MakeAnimalSpeak(animal Animal) {
	fmt.Println(animal.Speak())
}

func processAnimalTypeAssertion(animal Animal) {
	if dog, ok := animal.(*Dog); ok {
		fmt.Printf("Type: %T, Value: %v\n", dog, dog)
		fmt.Println(dog.Bark())
	}

	if cat, ok := animal.(*Cat); ok {
		fmt.Printf("Type: %T, Value: %v\n", cat, cat)
		fmt.Println(cat.Purr())
	}
}

func processAnimalTypeSwitch(animal Animal) {
	switch v := animal.(type) {
	case *Dog:
		fmt.Printf("Type: %T, Value: %#v\n", v, v)
		fmt.Println(v.Bark())
	case *Cat:
		fmt.Printf("Type: %T, Value: %#v\n", v, v)
		fmt.Println(v.Purr())
	default:
		fmt.Printf("Type: %T, Value: %#v\n", v, v)
	}
}

func main() {

	var animal Animal

	if animal != nil {
		fmt.Println("Animal is not nil")
	}

	dog := &Dog{}
	animal = dog
	fmt.Println(animal.Speak())
	fmt.Printf("Value animal %v, type animal %T\n", animal, animal)

	cat := &Cat{Name: "Тайсон"}

	if animal != nil {
		fmt.Println("Animal is not nil")
	}

	dog.Name = "Шайтан"
	fmt.Printf("Value animal %v, type animal %T\n", animal, animal)

	var ptr *struct{}
	var iface interface{}
	iface = ptr
	if iface != nil {
		println("iface is not nil")
	}

	var empty interface{}
	empty = dog
	fmt.Printf("Value empty %v, type empty %T\n", empty, empty)
	fmt.Println("empty name", empty.(*Dog).Name)
	fmt.Println("empty speak", empty.(*Dog).Speak())

	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)

	processAnimalTypeAssertion(dog)
	processAnimalTypeAssertion(cat)

	processAnimalTypeSwitch(dog)
	processAnimalTypeSwitch(cat)

}
