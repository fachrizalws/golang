package main
import (
    "fmt"
)
func main(){
    demoPanic()
}
func demoPanic() {
    fmt.Println("----contoh demo error handling---")
    defer func() {
        fmt.Println("mengerjakan sesuatu")
    }()
    panic("ini merupakan panic dari demoPanic()")
    fmt.Println("Pesan ini tidak akan muncul")
}
