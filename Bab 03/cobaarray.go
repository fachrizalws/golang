package main
import (
	"fmt"
	"math/rand"
	)
func main() {
	// mendefinisikan array
	fmt.Println("mendefinisikan arrays")
	var angka[6] int
	var kota[6] string
	var matrix[4][4] int // array 2D
	// memasukkan data
	fmt.Println(">>>>>memasukkan data array")
	for i:=0;i<6;i++ {
		angka[i] = i
		kota[i] = fmt.Sprintf("Kota %d",i)
		}
	// memasukkan data matrix
	fmt.Println(">>>>>memasukkan data matrix")
	for i:=0;i<4;i++ {
		for j:=0;j<4;j++ {
		matrix[i][j] = rand.Intn(100) //memasukkan data random interval sampai 100
        }
    }
	// menampilkan data
	fmt.Println(">>>>>Menampilkan data array")
	for j:=0;j<6;j++ {
		fmt.Println(angka[j])
		fmt.Println(kota[j])
	}
	// menampilkan data matrix
	fmt.Println(">>>>>menampilkan data matrix")
	for i:=0;i<4;i++ {
		for j:=0;j<4;j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
