package main
import “fmt”
func main() {
	dipilih := 2
	switch dipilih {
		case 0:
			fmt.Println(“dipilih = 0”)
		case 1:
			fmt.Println(“dipilih = 1”)
		case 2:
			fmt.Println(“dipilih = 2”)
		case 3:
			fmt.Println(“dipilih = 3”)
		default:
			fmt.Println(“lainnya..”)
	}
}
