package main
import "fmt"

//type one
var (
	name string
	age	 string
	location string
)

//type two

var (
	name,location string
	age	int
)

//type three

var name string 
var age int
var location string

//Variable Initialization

//type one
var (
	name string = "Ahmad"
	age	int = 20
	location string = "Cilacap"
)

//type two

var (
	name 	 = "Ahmad"
	age  	 = 20
	location = "Cilacap"
)

//type three

var (
	name,location,age = "Ahmad","Cilacap",20
)

//inside function

/*
func main(){
	name,location := "Ahmad","Cilacap"
	age := 20
	fmt.Printf("%s age %d from %s",name,age,location)
}
*/

//varible can contain any type, example is include function

func main(){
	action := func(){
		//ini actionnya
	}

	action() //dipanggil variablenya
}
