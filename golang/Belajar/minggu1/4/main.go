package main

import "fmt"

func sapa() {
	fmt.Println("halo, golang!")
}

func tambah(a int, b int) int {
	return a + b
}

func hitung(a, b int) (int, int){
	return a + b, a * b
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func ubahNilai(a *int) {
	*a = *a + 10
}

// struct dan method

type User struct{
	FirstName string
	LastName string
	Age int
}

type Userr struct{
	FirstNama string
	LastNama string
}

// method =>
func (p Userr) FullNama() string {
return p.FirstNama + " " + p.LastNama
}


// Pointer Receiver

type Pengguna struct{
	Nama string
	Umur int
}

func (s *Pengguna) TambahUmur() {
	s.Umur++
}

func main() {
	sapa()

	hasil := tambah(5, 3)
	fmt.Println("Hasil Penjumlahan:", hasil)

	jumlah, kali := hitung(3, 5)
	fmt.Println("Jumlah:", jumlah)
	fmt.Println("kali:", kali)

	fmt.Println(sum(1,2,3))
	fmt.Println(sum(4,5,6,7,8))


	fmt.Println("awal")
	defer fmt.Println("ditunda hingga akhir")
	fmt.Println("Akhir")

// pointer

	x :=5 
	fmt.Println("sebelum: ",x)

	ubahNilai(&x)
	fmt.Println("sesudah:", x)

u := User{FirstName: "Didi", LastName: "Sutisno", Age: 22}
fmt.Println("Nama:", u.FirstName, u.LastName)
fmt.Println("umur:", u.Age)

user := Userr{"Andi", "saputro"}
    fmt.Println("Nama lengkap:", user.FullNama())

s := Pengguna{Nama: "Citra", Umur: 20}
fmt.Println("Umur Sebelum:", s.Umur)

s.TambahUmur()
fmt.Println("Umur Sesudah:", s.Umur)

}