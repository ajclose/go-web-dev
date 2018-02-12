package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println("Hello, my name is", p.fname, p.lname)
}

func (sa secretAgent) saSpeak() {
	fmt.Println("I have the license to kill", sa.licenseToKill)
}

func main() {
	p1 := person{"Al", "Jon"}
	p1.pSpeak()
	sa1 := secretAgent{person{"James", "Bond"}, true}
	sa1.person.pSpeak()
	sa1.saSpeak()
}
