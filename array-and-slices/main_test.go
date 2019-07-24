package main

import "testing"

//TestSumArray testing using array
func TestSumArray(t *testing.T){
	assertCorrectMessage := func(t *testing.T, got, want int) {
        t.Helper()
        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
	}
	
	t.Run("Adding numbers inside array and returning sum", func(t *testing.T){
		arraynumbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := SumArray(arraynumbers)
        want := 55

		assertCorrectMessage(t, got, want)
	})
}


//TestSumSlice testing using array
func TestSumSlice(t *testing.T){
	assertCorrectMessage := func(t *testing.T, got, want int) {
        t.Helper()
        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
	}
	
	t.Run("Adding numbers inside array and returning sum", func(t *testing.T){
		slicenumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := SumSlice(slicenumbers)
        want := 55

		assertCorrectMessage(t, got, want)
	})
}
