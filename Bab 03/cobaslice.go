package main
import (
	"fmt"
	"math/rand"
)
func main() {
	// mendefinisikan slice
	fmt.Println("mendefinisikan slices")
	var numbers[] int
	numbers = make([]int,6)
	matrix := make([][]int,4*4)
	// memasukkan data
	fmt.Println(">>>>>memasukkan data slice ")
	for i:=0;i<6;i++ {
		numbers[i] = rand.Intn(100) //masukan random
    }
	//memasukkan data matrix 
	fmt.Println(">>>>>memasukkan data slice matrix")
	for i:=0;i<4;i++ {
		matrix[i] = make([]int,4)
		for j:=0;j<4;j++ {
			matrix[i][j] = rand.Intn(100)
			}
	}
	// menampilkan data
	fmt.Println(">>>>>menampilkan data sclice")
	for j:=0;j<6;j++ {
		fmt.Println(numbers[j])
	}
	//menampilkan data matrix
	fmt.Println(">>>>>menampilkan data sclice dua dimensi")
	for i:=0;i<4;i++ {
		for j:=0;j<4;j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
