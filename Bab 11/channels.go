package main
import (
	"fmt"
)
func main(){
	fmt.Println("simple channel")
	// mendefinisikan sebuah channel
	c := make(chan int)
	// menjalankan sebuah fungsi di background
	go func() {
		fmt.Println("memproses goroutine")
		c <- 10   // menulis data untuk sebuah channel
	}()
	val := <-c // membaca data untuk sebuah channel
	fmt.Printf("nilai: %d\n",val)
}
