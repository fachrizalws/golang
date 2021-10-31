package main
import (
	"fmt"
	"log"
)
func main() {
	simpleLogging()
}
func simpleLogging(){
	fmt.Println("-----------logging sederhana-----------")
	log.Println("Halo Dunia")
	log.Println("Ini merupakan contoh pesan error sederhana")
}
