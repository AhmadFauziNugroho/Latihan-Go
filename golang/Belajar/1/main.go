package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// variable => var atau :=

	// cara panjang
	var nama1 string = "Sienaf"

	// cara singkat
	nama2 := "Sienaf"

	// banyak sekaligus
	var (
		umur  int     = 20
		berat float64 = 55.5
		aktif bool    = true
	)

	// konstanta = Nilai Tidak Berubah => const
	const phi = 3.14
	const appName = "InvenAPPs"

	fmt.Println(nama1, nama2, umur, berat, aktif, phi, appName)
}
