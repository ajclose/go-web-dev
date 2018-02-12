package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func main() {
	p1 := person{"James", "Bond", []string{"olives", "vodka", "beef"}}
	fmt.Println(p1.favFood)
	for _, f := range p1.favFood {
		fmt.Println(f)
	}
}
