package main
import “fmt”
func main() {
    var (
        x = 6
        y = 26
    )
    fmt.Println(x>y && x!=y)
    fmt.Println(!(x>=y))
    fmt.Println(x==y || x>y)
}
