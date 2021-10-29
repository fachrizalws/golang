package main
import "fmt"
func main(){
	var x int
	x = 26
	fmt.Println(x)  // menampilkan nilai dari x
	fmt.Println(&x) // menampilkan lokasi dari x
	// mendefinisikan pointer
	var num *int
	val := new(int)
	num = new(int)
	*num = x   //menentukan nilai
	val = &x   //menentukan lokasi
	fmt.Println("===pointer num===")
	fmt.Println(*num)		// menampilkan nilai dari x
	fmt.Println(num) // menampilkan lokasi dari x
	fmt.Println("===pointer val===")
	fmt.Println(*val)		//menampilkan nilai dari x
	fmt.Println(val) //menampilkan lokasi dari x
}
