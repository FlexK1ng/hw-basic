package main

import "fmt"

type User struct {
	Age  int
	Data int
	Name string
}

func main() {
	a := &User{Age: 1, Data: 20, Name: "Bob"}
	b := a
	b.Age = 12
	YaNeponyal(*a)
	fmt.Println(a)
	fmt.Println(*b)

}

func YaNeponyal(c User) {
	c.Age = 24
	fmt.Println(c)

}
