package main
import (
    "fmt"
    "log"
	"os"
)
func main() {
    formattingLogging()
}
func formattingLogging(){
    fmt.Println("-----------formattingLogging-----------")
    var warning *log.Logger
    warning = log.New(
        os.Stdout,
        "PERINGATAN: ",
        log.Ldate|log.Ltime|log.Lshortfile)
    warning.Println("Pesan peringatan pertama")
    warning.Println("Pesan peringatan kedua")
}
