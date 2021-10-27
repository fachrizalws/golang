package main
import “fmt”
func main() {
	// iterasi perintah for
	var i int
	for i=0;i<6;i++ {
	if i==2 {
		break
		}
		fmt.Println(i)
	}
	for j:=6;j< 12;j++ {
		if j==9 {
			continue
		}
		fmt.Println(j)
	}
}
