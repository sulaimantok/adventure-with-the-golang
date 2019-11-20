package main

import "fmt"

func sum(a []int,c chan int){
	sum := 0
	for _, v := range a{
		sum += v
	}
	//send sum to c
	c <- sum 
}

func main(){
	a := []int{7,2,8,-9,4,0}

	//membuat channel
	c := make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	//receive from c
	x,y := <-c,<-c


	fmt.Println(x,y,x+y)
}