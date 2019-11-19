package main
import "fmt"

//function adding x and y
func add(x int, y int){
	return x+y
}

/* this same with the function before this
function but, this define the data type of return value
func add(x int, y int) int {
	return x+y
}

*/

//function reduction x and y
func minus(x int, y int){
	return x-y
}


//Main Function
func main(){
	fmt.Println("Adding Output 50 and 22 : %d",add(50,22))
	fmt.Println("Reduction Output 50 and 22 : %d",minus(50,22))
}