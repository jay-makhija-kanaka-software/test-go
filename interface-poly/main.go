package main

import (
	"fmt" 
"github.com/GoesToEleven/learn-to-code-go-version-03"
)

type person struct {
	name string
}

type secretAgent struct {
	person
	canKill bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Hello my name is ", p.name)
}

func humanSay(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Jay",
	}
	p1.speak()

	sa1 := secretAgent{
		person: person{
			name: "James Bond",
		},
		canKill: true,
	}

	sa1.speak()

	humanSay(p1)
	humanSay(sa1)
}
