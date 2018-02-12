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

type human interface {
	speak() string
}

func (p person) speak() string {
	return fmt.Sprintln("Hello, my name is", p.fname, p.lname)
}

func (sa secretAgent) speak() string {
	return fmt.Sprintln("I have the license to kill", sa.licenseToKill)
}

func talk(h human) {
	fmt.Println(h.speak())
}
func main() {
	p1 := person{"Al", "Jon"}
	talk(p1)
	sa1 := secretAgent{person{"James", "Bond"}, true}
	talk(sa1)
}
