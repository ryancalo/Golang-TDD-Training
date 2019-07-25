package main

import "fmt"

type Rectangle struct {
	Width float64
    Height float64
}

type Geometry interface {
	PerimeterMethod() float64
	AreaMethod() float64
}

func Measure(g Geometry) (float64, float64){
  return g.AreaMethod(), g.PerimeterMethod()
}


func Perimeter(rectangle Rectangle)float64{
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
    return rectangle.Width * rectangle.Height
}


func (r Rectangle) PerimeterMethod() float64{
	return 2 * (r.Width + r.Height)
}


func (r Rectangle) AreaMethod() float64{
	return r.Width * r.Height
}

func main(){
	rectangle := Rectangle{
						     Width: 48.0,
							 Height: 24.0,
						   }
	perim := Perimeter(rectangle)
	area := Area(rectangle)
	fmt.Printf("The Perimeter of the Rectangle is: %.2f \n", perim)
	fmt.Printf("and the Area is: %.2f \n", area)

	perimMethod := rectangle.PerimeterMethod()
	areaMethod := rectangle.AreaMethod()
	fmt.Printf("The Perimeter of the Rectangle using method is: %.2f \n", perimMethod)
	fmt.Printf("and the Area using Method is: %.2f \n", areaMethod)


	//Using Interface
	iarea, iperim := Measure(rectangle)
	fmt.Printf("The area of the rectangle using a interface is: %.2f with a perimeter of: %.2f \n", iarea,iperim)

}