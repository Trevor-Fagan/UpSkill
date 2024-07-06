package main

import (
	"fmt"
	"log"
)

type makeNoise interface {
	speak() string
}

type person struct {
	name string
	age int
	gender rune
}

type dog struct {
	name string
	age int
	gender rune
}

func main () {
	trevor := person{name: "Trevor", age: 22, gender: 'M'}
	brownie := dog{name: "Brownie", age: 13, gender: 'F'}

	saySomething(&trevor)
	saySomething(&brownie)
}

func saySomething(t makeNoise) {
	fmt.Println(t.speak())
}

func (p *person) speak() string {
	return "Hello!"
}

func (p *dog) speak() string {
	return "Woof!"
}