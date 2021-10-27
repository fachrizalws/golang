package main
import “fmt”
func main() {
    // mendeklarasikan variabel-variabel
    var str string
    var n, m int
    var mn float32
    // menetapkan variabel
    str = “Halo, Go!!!”
    n = 20
    m = 60
    mn = 2.66
    fmt.Println(“nilai dari str= “,str)
    fmt.Println(“nilai dari n= “,n)
    fmt.Println(“nilai dari m= “,m)
    fmt.Println(“nilai dari mn= “,mn)
    // mendeklarasikan dan menetapkan nilai untuk  tiap variabel
    var kota string = “Semarang”
    var x int = 600
    fmt.Println(“Nilai dari kota= “,kota)
    fmt.Println(“Nilai dari x= “,x)
    // mendeklarasikan dengan mendefinisikan tipe data
    negara := “ID”
    nilai := 17
    fmt.Println(“isian dari negara= “,negara)
    fmt.Println(“isian dari nilai= “,nilai)
    // mendefinisikan beberapa variabel
    var (
        nama string
        email string
        umur int
    )
    nama = “Fachrizal Rian Pratama”
    email = “fachrizal.rian.p@walisongo.ac.id”
    umur = 17
    fmt.Println(nama)
    fmt.Println(email)
    fmt.Println(umur)
}
