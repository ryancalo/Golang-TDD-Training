package main

import "fmt"

func Add(x,y int) int{
	return x + y
}

func main(){
	sum := Add(3,4)
	fmt.Println(sum)
}