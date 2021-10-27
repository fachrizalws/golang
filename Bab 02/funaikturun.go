package main
import “fmt”
func main() {
	// mendeskripsikan variabel
	var x = 4
	// fungsi naik
	fmt.Printf(“x = %d \n”,x)
	x = x + 1
	fmt.Printf(“x + 1 = %d \n”,x)
	x++
	fmt.Printf(“x++ = %d \n”,x)
	// fungsi turun
	x = x - 1
	fmt.Printf(“x - 1 = %d \n”,x)
	x--
	fmt.Printf(“x--=%d \n”,x)
}
