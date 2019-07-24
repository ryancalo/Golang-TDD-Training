package main

import "fmt"

func Repeat(char string)string{
	var repeated string
	
	for i := 1; i < 30; i++ {
		repeated = repeated + char
	}
	return repeated
}

func main(){
	//i put z kay katugon nako
	result := Repeat("z")
	fmt.Println(result)
}