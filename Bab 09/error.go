package main
import (
    "fmt"
)
func main(){
    calculate()
}
func calculate() {
    fmt.Println("----contoh demo error handling---")
    x := 10
    y := 0
    z := 0
    z = x/y
    fmt.Printf("Hasil = %.2f \n",z)
    fmt.Println("Selesai")
}
