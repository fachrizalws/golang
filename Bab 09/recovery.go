package main
import (
    "fmt"
)
func main(){
    calculate2()
}
func calculate2() {
    fmt.Println("----contoh demo error handling---")
    z := 0
    defer func() {
        x := 10
        y := 0
        z = x/y
        if error := recover(); error != nil {
            fmt.Println("Menanggapi....", error)
            fmt.Println(error)
        }
    }()
    fmt.Printf("hasil = %d \n",z)
    fmt.Println("Selesai")
}
