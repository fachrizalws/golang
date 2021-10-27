package main
import (
    “fmt”
    “math”
)
func main(){
    x := 2.6
    y := 1.9
    z := math.Pow(x,2)
    fmt.Printf(“%.2f^%d = %.2f \n”,x,2,z)
    z = math.Sin(x)
    fmt.Printf(“Sin(%.2f) = %.2f \n”,x,z)
    z = math.Cos(y)
    fmt.Printf(“Cos(%.2f) = %.2f \n”,y,z)
    z = math.Sqrt(x*y)
    fmt.Printf(“Sqrt(%.2f*%.2f) = %.2f \n”,x,y,z)
}
