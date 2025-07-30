package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil parameter 'name' dari URL query 
	name := r.URL.Query().Get("name")
	if name == "" { // Jika 'name' kosong, gunakan 'Guest' 
		name = "Guest"
	}
	// Menampilkan pesan selamat datang 
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Mendaftarkan handler untuk rute "/hello" 
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	// Menjalankan server di port 8080 
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}