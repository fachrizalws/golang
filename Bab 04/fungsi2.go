package main
import (
	"fmt"
	"math"
)
func main(){
	foo()
	luas_lingkaran(2.56)
	kalkulasi(2, 6.7, 5)
	val := kalkulasi_lanjutan(2, 4.8, 7)
	fmt.Printf("kalkulasi_lanjutan() = %.2f \n",val)
	val1,val2,val3 := hitung(6, 2.7, 1.4, "energi")
	fmt.Printf("val1=%.2f, val2=%.2f, val3=\"%s\" \n", val1, val2, val3)
	hasil := add(1,2,3,4,5,6,7,8,9,10,11)
	fmt.Printf("add() = %d \n",hasil)
	fungsi_penutup("percobaan")
	ret := fibonacci(15)
	fmt.Printf("fibonacci() = %d \n",ret)
}

//contoh fungsi sederhana
func foo() {
	fmt.Println("foo() telah dipanggil")
}

//contoh fungsi dengan satu parameter
func luas_lingkaran(r float64){
	luas := math.Pi * math.Pow(r,2)
	fmt.Printf("Luas lingkaran (r=%.2f) = %.2f \n",r, luas)
}

//Contoh fungsi dengan beberapa parameter
func kalkulasi(a,b, c float64){
	hasil := a*b*c
	fmt.Printf("a=%.2f, b=%.2f, c=%.2f = %.2f \n",a, b, c, hasil)
}
//Contoh fungsi dengan parameter dan nilai kembali
func kalkulasi_lanjutan(a,b, c float64) float64 {
	hasil := a*b*c
	return hasil
}
//Contoh fungsi dengan parameter dan beberapa nilai kembali
func hitung(a,b, c float64, nama string) (float64,float64,string){
	hasil1 := a*b*c
	hasil2 := a + b + c
	hasil3 := hasil1 + hasil2
	namabaru := fmt.Sprintf("%s nilai = %.2f", nama,hasil3)
	return hasil1, hasil2, namabaru
}
//Contoh fungsi dengan nol/lebih parameter & nilai kembali
func add(numbers ...int) int {
	hasil := 0
	for _, number := range numbers {
		hasil += number
	}
	return hasil
}
//Contoh fungsi penutup
func fungsi_penutup(nama string){
	hoo := func(a,b int){
		hasil := a*b
		fmt.Printf("hoo() = %d \n",hasil)
	}
	joo := func(a,b int) int {
		return a*b + a
	}
	fmt.Printf("fungsi_penutup(%s) telah dipanggil\n",nama)
	hoo(2,3)
	val := joo(5,8)
	fmt.Printf("val from joo() = %d \n",val)
}
//Contoh fungsi rekursi
func  fibonacci(n int) int {
	if n==0{
		return 0
	}else if n==1 {
		return 1
	}
	return (fibonacci(n-1) + fibonacci(n-2))
}
