package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	jimmy := person{firstName: "jimmy", lastName: "korn", contact: contactInfo{email: "jimmykorn@fuck.com", zipcode: 696969}}

	jimmy.updateName("johnson")
	jimmy.print()
}

func (p person) updateName(newFirstName string){
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}