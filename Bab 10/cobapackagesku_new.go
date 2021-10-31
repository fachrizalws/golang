package main
import (
    "fmt"
    "packagesku"
)
func main(){
    fmt.Println("mengakses packagesku...")
    var c int
    c = packagesku.Add(5,8)
    fmt.Printf("add()=%d\n",c)
    c = packagesku.Subtract(5,8)
    fmt.Printf("subtract()=%d\n",c)
    c = packagesku.Multiply(5,3)
    fmt.Printf("multiply()=%d\n",c)
}
