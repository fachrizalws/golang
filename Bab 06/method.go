package main
import (
    "fmt"
    "math"
)
//mendefinisikan sebuah struct
type Lingkaran struct {
    x,y int
    r float64
}
func (c Lingkaran) display() {
    fmt.Printf("x=%d,y=%d,r=%.2f\n",c.x,c.y,c.r)
}
func (c Lingkaran) area() float64 {
    return math.Pi * math.Pow(c.r,2)
}
func (c Lingkaran) circumference() float64 {
    return 2* math.Pi * c.r
}
func (c Lingkaran) moveTo(newX,newY int) {
    c.x = newX
    c.y = newY
}
func main() {
	//methods
	shape := Lingkaran{10,5,2.8}
	shape.display()
	fmt.Printf("area=%2.f\n",shape.area())
	fmt.Printf("circumference=%2.f\n",shape.circumference())
	shape.moveTo(5,5)
	shape.display()
}
