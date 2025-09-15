package main

import (
	"5/helper"
	"errors"
	"fmt"
)

func main() {
    fmt.Println("Memanggil package helper:")
    helper.CetakHello()

    hasil, err:= helper.Bagi(10, 2)
    if err != nil {
        fmt.Println("Terjadi error:", err)
        return
    }
    fmt.Println("hasil:", hasil)

    if err := helper.CekUmur(19); err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Pendaftaran Berhasil")
    }

    err = helper.Proses([]int{})
    if errors.Is(err, helper.ErrDataKosong) {
        fmt.Println("tidak ada data yang bisa diproses") 
        return
    }
    fmt.Println("Proses Selesai")
}
