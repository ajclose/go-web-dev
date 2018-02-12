package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

func main() {
	p1 := person{"James", "Bond", []string{"olives", "vodka", "beef"}}
	s := p1.walk()
	fmt.Println(s)
}
