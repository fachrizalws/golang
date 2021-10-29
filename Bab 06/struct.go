package main
import (
    "fmt"
    "time"
)

//mendefinisikan sebuah struct
type Mahasiswa struct {
    nim int
    nama string
    jurusan string
    waktu time.Time
}

func main() {
    //mendefinisikan variabel
    var mhs Mahasiswa
    newMhs := new(Mahasiswa)
 	
	  //menentukan nilai
    mhs.nim = 2
    mhs.nama = "Mahasiswa 2"
    mhs.jurusan = "Fisika"
    mhs.waktu = time.Now()
    newMhs.nim = 5
    newMhs.nama = "Mahasiswa 5"
    newMhs.jurusan = "Biologi"
    newMhs.waktu = time.Now()
   	
	  //menampilkan
    fmt.Println("=====================")
    fmt.Println(mhs.nim)
    fmt.Println(mhs.nama)
    fmt.Println(mhs.jurusan)
    fmt.Println(mhs.waktu)
    fmt.Println("=====================")
    fmt.Println(newMhs.nim)
    fmt.Println(newMhs.nama)
    fmt.Println(newMhs.jurusan)
    fmt.Println(newMhs.waktu)
}
