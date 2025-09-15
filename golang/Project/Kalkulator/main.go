package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Masukan Angka Pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukan Angka Kedua: ")
	fmt.Scan(&b)

	fmt.Println("Hasil Penjumlahan: ", a+b)
	fmt.Println("Hasil Pengurangan: ", a-b)
	fmt.Println("Hasil Perkalian: ", a*b)
	if b != 0{
		fmt.Println("Hasil Pembagian: ",a/b)
	}else {
		fmt.Println("Tidak Bisa Dibagi Dengan nol")
	}

}