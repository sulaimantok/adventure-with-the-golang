package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string { //Name method used for type User
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string { //Name method used for type Customer
	return c.FullName
}

type Namer interface {
  Name() string //Both Name() methods can be called using the Namer interface
}

func Greet(n Namer) string {
	return fmt.Sprintf("Hello %s", n.Name())
}

func main() {
	u := &User{"Ahmad", "Kamal"}
	fmt.Println(Greet(u))
	c := &Customer{62, "Indonesia"}
	fmt.Println(Greet(c))
}