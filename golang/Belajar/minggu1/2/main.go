package main

import "fmt"

func main() {
	var nama string
	fmt.Print("masukan Nama Kamu: ")
	fmt.Scan(&nama)
	fmt.Println("Halo,", nama)

	a := 10
	b := 3

	fmt.Println("Penjumlahan: ", a+b)
	fmt.Println("Pengurangan: ", a-b)
	fmt.Println("Perkalian: ", a*b)
	fmt.Println("Pembagian: ", a/b)
	fmt.Println("Sisa Bagi:", a%b)

	fmt.Println(a == b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	
	x := true
	y := false
	fmt.Println(x && y)
	fmt.Println(x || y)
	fmt.Println(!x)
}