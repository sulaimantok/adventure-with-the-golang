package main

import "fmt"

func main(){
	//array with string value
	var a [2]string
	// a := [2]string{"Hello","World"}
	// a := [...]

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	//array with integer value
	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}