package main
import "fmt"

//Constants

const Pi = 3.14

const (
	StatusOK					= 200
	StatusCreated				= 201
	StatusAccepted				= 202
	StatusNonAuthoritativeInfo	= 203
	StatusNoContent				= 204
	StatusResetContent			= 205
	StatusPartialContent		= 206
)

//Printing

func main(){
	cylonModel := 6
	fmt.Printf(Pi)
	fmt.Println(cylonModel)
}

