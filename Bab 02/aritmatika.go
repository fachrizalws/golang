package main
import “fmt”
func main() {
    // mendeskripsikan variabel-variabel
    var x, y int
    // menentukan nilai
    x = 2
    y = 6
    // operasi aritmatika
    // penambahan
    z := x + y
    fmt.Printf(“%d + %d = %d \n”,x,y,z)
    // pengurangan
    w := x - y
    fmt.Printf(“%d - %d = %d \n”,x,y,w)
    // pembagian
    v := float32(x) / float32(y)
    fmt.Printf(“%d / %d = %.2f \n”,x,y,v)
    // perkalian
    u := x * y
    fmt.Printf(“%d * %d = %d \n”,x,y,u)
}
