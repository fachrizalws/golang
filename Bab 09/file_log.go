package main
import (
    "fmt"
    "log"
	"os"
)
func main() {
    fileLogging()
}
func fileLogging(){
  fmt.Println("-----------file logging-----------")
	file, err := os.OpenFile("C:/Users/TOSHIBA/logging/error.log",
	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("Gagal membuka file log.")
        return
    }
    var logFile *log.Logger
    logFile = log.New(
        file,
        "APP: ",
        log.Ldate|log.Ltime|log.Lshortfile)
    logFile.Println("Pesan error pertama")
    logFile.Println("Pesan error kedua")
    logFile.Println("Pesan error ketiga")
    fmt.Println("Selesai")
}
