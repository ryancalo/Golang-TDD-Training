package main

import "fmt"

func SumArray(numbers [10]int) int{
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlice(numbers []int)int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main(){
	arraynumbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slicenumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	asum := SumArray(arraynumbers) 
    ssum := SumSlice(slicenumbers)
	fmt.Printf("Sum using array : %d \n", asum)
	fmt.Printf("Sum using slice : %d \n", ssum)
}