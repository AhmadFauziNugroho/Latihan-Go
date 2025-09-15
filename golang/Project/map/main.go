package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukan Kalimat: ")
	kalimat, _ := reader.ReadString('\n')

	kalimat = strings.TrimSpace(kalimat)

	kata := strings.Fields(kalimat)

	frekuensi := make(map[string]int)
	for _, k := range kata {
		frekuensi[k]++
	}

	fmt.Println("\nFrekuensi kata:")
	for k, v := range frekuensi {
		fmt.Printf("%s: %d\n", k, v)
	}
}
