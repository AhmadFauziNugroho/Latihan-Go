package helper

import (
    "fmt"
    "errors"
)

func CetakHello() {
    fmt.Println("Halo dari package helper!")
}

func Bagi(a,b int) (int, error) {
    if b == 0 {
        return 0, errors.New("Tidak bisa dibagi dengan 0")
    }
    return a / b, nil
}

func CekUmur(umur int) error{
    if umur < 18 {
        return fmt.Errorf("umur %d terlalu muda untuk mendaftar", umur)
    }
    return nil
}

var ErrDataKosong = errors.New("data kosong")

func Proses(data []int) error {
    if len(data) == 0 {
        return ErrDataKosong
    }
    return nil
}
