package main

import "fmt"

func main() {

// if else
	
	angka := 15

	if angka > 10 {
		fmt.Println("angka lebih besar dari 10")
	} else if angka == 10 {
		fmt.Println("angka sama dengan 10")
	} else {
		fmt.Println("angka lebih kecil dari 10")
	}

// switch

	hari := "senin"

	switch hari {
	case "senin":
		fmt.Println("hari pertama kerja")
	case "jumat":
		fmt.Println("hari terakhir kerja")
	default:
		fmt.Println("hari biasa")
	}

// for loop

	for i := 1; i <= 5; i++ {
		fmt.Println("Iterasi ke: ", i)
	}

	nomor := 1
	for nomor < 5 {
		fmt.Println("nomor: ", nomor)
		nomor++
	}

	buah := []string{"apel","jeruk","anggur"}
	for i, b := range buah {
		fmt.Printf("index %d: %s\n", i, b)
	}

// array

	var number [3]int
	number[0] = 10
	number[1] = 12
	number[2] = 14

	fmt.Println("isi array:", number)
	fmt.Println("panjang array:", len(number))

// slice

numero := []int{10,20,30}
fmt.Println("slice awal:",numero)

numero = append(numero, 40)
fmt.Println("setelah append:",numero)

bagian := numero[1:3]
fmt.Println("sub-slice:", bagian)

fmt.Println("panjang:", len(numero))
fmt.Println("kapasitas:", cap(numero))

// map

umur := map[string]int{
	"uji":22,
	"someone":20,
}
fmt.Println("umur Uji:", umur["uji"])

umur["tia"] = 23

delete(umur, "someone")

for nama, u := range umur {
	fmt.Printf("%s berubur %d\n", nama, u)
}
}