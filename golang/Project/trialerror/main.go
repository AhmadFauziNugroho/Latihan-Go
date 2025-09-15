package main

import (
	"fmt"
	"Project/trialerror/util"
)

func main() {
	angka, err := util.MustPositive(5)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("angka valid", angka)
}