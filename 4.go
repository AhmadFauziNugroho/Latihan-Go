package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Request receive: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini adalah endpoint GET /todos. Daftar Tugas akan ditampilkan."))
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini adalah endpoint Post /todos. Daftar Tugas akan ditambahkan."))
}

func main() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	r.HandleFunc("/todos", getTodosHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")

	fmt.Println("Server berjalan di http://localhost:8080/todos")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}