package main
import (
    "fmt"
    "math/rand"
)
func main() {
    //mendefinisikan array
    fmt.Println("mendefinisikan map")
    produk := make(map[string]int)
    pelanggan := make(map[string]int)
    //memasukkan data
    fmt.Println(">>>>>memasukkan data map")
    produk["produk1"] = rand.Intn(100)
    produk["produk2"] = rand.Intn(100)
    pelanggan["pelanggan1"] = rand.Intn(100)
    pelanggan["pelanggan2"] = rand.Intn(100)
		//menampilkan data
    fmt.Println(">>>>>menampilkan data map")
    fmt.Println(produk["product1"])
    fmt.Println(produk["product2"])
    fmt.Println(pelanggan["pelanggan1"])
    fmt.Println(pelanggan["pelanggan2"])
}
