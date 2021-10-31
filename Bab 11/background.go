package main
import (
	"fmt"
		"time"
)
func main(){
	fmt.Printf("contoh goroutines\n")
	// fungsi berikut berjalan di background
	go hitung()
	index := 0
	// berjalan di background
	go func() {
		for index < 6 {
			fmt.Printf("go func() index= %d \n", index)
			var hasil float64 = 2.5 * float64(index)
			fmt.Printf("go func() hasil= %.2f \n", hasil)
			time.Sleep(500 * time.Millisecond)
			index++
		}
	}()
	// berjalan di background
	go fmt.Printf("menampilkan go yang berjalan di background\n")
	// Silahkan tekan ENTER untuk keluar
	var input string
	fmt.Scanln(&input)
	fmt.Println("selesai")
}
func hitung() {
	i := 12
	for i < 15 {
		fmt.Printf("hitung()= %d \n", i)
		var hasil float64 = 8.2 * float64(i)
		fmt.Printf("hitung() hasil= %.2f \n", hasil)
		time.Sleep(500 * time.Millisecond)
		i++
	}
}
