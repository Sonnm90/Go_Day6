package _defer

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func Demo() {
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}
func printA(a *int) {
	fmt.Println(*a)
	*a = 8
	fmt.Println("value of a in deferred function", *a)
}

func Test() {
	a := 5
	defer printA(&a)
	//a++
	fmt.Println("value of a before deferred function call", a)

}
