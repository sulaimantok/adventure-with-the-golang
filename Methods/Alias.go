package main

import (
	"fmt"
	"math"
	"string"
)
//Alias
type stringku string 
type floatku float64

func (s stringku) Uppercase() string{
	return strings.ToUpper(string(s))	
}

func (f floatku) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main(){
	fmt.Println(stringku("Test").Uppercase())
	f := floatku(-math.Sqrt2)
	fmt.Println(f.Abs)
}