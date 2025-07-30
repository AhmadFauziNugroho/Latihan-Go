package main

import(
	"fmt"
)

func main() {
	var num1, num2 float64

	fmt.Print("masukan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Print("masukan angka kedua: ")
	fmt.Scanln(&num2)

	fmt.Println("Penjumlahan:", num1+num2)
	fmt.Println("Pengurangan:", num1-num2)
	fmt.Println("Perkalian:",num1*num2)
	if num2 != 0 {
		fmt.Println("Pembagian:", num1/num2)
	} else {
		fmt.Println("Pembagian: Tidak dapat dibagi nol")
	}
}
