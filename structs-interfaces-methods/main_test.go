package main

import "testing"


func TestPerimeter(t *testing.T){
	rectangle := Rectangle{
		Width: 48.0,
		Height: 24.0,
	  }


   assertCorrectMessage := func(t *testing.T, got, want float64) {
        t.Helper()
        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
	}
	
	t.Run("Getting perimeter of a rectangle", func(t *testing.T){
		got := Perimeter(rectangle)
        want := 144.00

		assertCorrectMessage(t, got, want)
	})
	
}


func TestArea(t *testing.T){
	rectangle := Rectangle{
		Width: 48.0,
		Height: 24.0,
	  }

	assertCorrectMessage := func(t *testing.T, iarea, iareaW, iperim, iperimW  float64) {
        t.Helper()
        if iarea != iareaW && iperim != iperimW {
            t.Errorf("i got %.2f in area but i want %.2f in Getting Area And i got %.2f in perim but i want %.2f in Getting Perimeter", iarea, iareaW, iperim, iperimW)
        }
	}

	
	t.Run("Getting Area of a rectangle", func(t *testing.T){
		iarea, iperim := Measure(rectangle)
        iareaW, iperimW :=  1152.00, 144.00

		assertCorrectMessage(t, iarea, iareaW, iperim, iperimW)
	})
	
}


func TestMeasure(t *testing.T){

}