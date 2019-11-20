package main

import(
	"fmt"
	"time"
)

type Errorku struct{
	When time.time
	What string
}

func (e *Errorku) Error() string{
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error{
	return &Errorku{
		time.Now(),
		"Ini tidak Bekerja",
	}
}

func main(){
	if err := run(); err != nil{
		fmt.Println(err)
	}
}