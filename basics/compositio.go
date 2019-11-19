package main

import "fmt"

type User struct {
	Id       int
	Name     string
	Location string
}

//type Player with one additional attribute

type Player struct {
	Id       int  
	Name     string
	Location string
	GameId	 int
}

func main() {
	p := Player{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p) // the value in a default format when printing structs,
                        // the plus flag (%+v) adds field names
}