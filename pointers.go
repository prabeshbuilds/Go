package main
import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
	Base   float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}


func main() {
	rect := Rectangle{Width: 5.0, Height: 3.0, Base: 4.0}
	area := rect.Area()
	fmt.Println("Area of the rectangle:", area)
}