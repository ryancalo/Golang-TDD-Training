package main

import "testing"

//TestHello Testing for simple Hello World
func TestHello(t *testing.T){

	assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
	}
	
	t.Run("Saying Hello Wolrld", func(t *testing.T){
		got := Hello()
        want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

}



//TestHelloArg testing for Hello with Argument
func TestHelloArg(t *testing.T){
  
	assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
	}
	
	t.Run("Saying Hello Wolrld", func(t *testing.T){
		got := HelloArg("Ryan")
        want := "Hello, Ryan"

		assertCorrectMessage(t, got, want)
	})

}