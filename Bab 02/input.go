package main
import (
	“fmt”
	“math”
)
func main() {
	//Menghitung Keliling Lingkaran
	fmt.Print(“Masukkan nilai jari-jari lingkaran: “)
	var jari2 float64
	fmt.Scanf(“%f”, &jari2)
	keliling := 2 * math.Pi * jari2
	fmt.Printf(“Keliling lingkaran dengan jari-jari %.2f = %.2f \n”,jari2, keliling)
	
	//Mengitung Luas Lingkaran
	luas := math.Pi * math.Pow(jari2,2)
	fmt.Printf(“Luas lingkaran dengan jari-jari %.2f = %.2f \n”,jari2, luas)
}
