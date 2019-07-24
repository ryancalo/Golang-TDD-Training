package main

import "fmt"

func Hello() string{
    return "Hello, World"
}

func HelloArg(name string) string{
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}



func main(){
 
	 fmt.Println(Hello())
	 fmt.Println(HelloArg("Ryan"))
	
}