package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Walk()
}

type Dog struct {
	Color string
	Address *Address
	M map[string]int
}

func (d Dog) Walk() {
	fmt.Println("Dog walks")
}

type RobotDog struct {
	Dog
	Brand string
}

type Address struct {
	Street string
}

func main() {
	var dog1, dog2 Animal
	aDog := &Dog{
		Color: "black",
		Address: &Address{
			Street: "124 Main Street",
		},
		M: map[string]int{
            "a": 1,
            "b": 2,
        },
	}

	bDog := &Dog{
		Color: "black",
		Address: &Address{
			Street: "124 Main Street",
		},
		M: map[string]int{
            "a": 1,
            "b": 2,
        },
	}

	dog1 = aDog
	dog2 = bDog

	fmt.Println(dog1)
	fmt.Println(dog2)
	fmt.Println(dog1 == dog2)
	fmt.Println(reflect.DeepEqual(dog1, dog2))

	//------------
	robotDog := &RobotDog{
		Dog: *aDog,
		Brand: "Mi",
	}

	robotDog.Walk()

	sl := []int{0, 1, 2, 3}

	fmt.Println(sl[:0])
}