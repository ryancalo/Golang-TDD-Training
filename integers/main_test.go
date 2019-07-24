package main

import "testing"

func TestAdd(t *testing.T){
	assertCorrectMessage := func(t *testing.T, got, want int) {
        t.Helper()
        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
	}
	
	t.Run("Adding Two integers and returning sum", func(t *testing.T){
		got := Add(10,25)
        want := 35

		assertCorrectMessage(t, got, want)
	})
}