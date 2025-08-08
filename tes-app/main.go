package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Food struct {
	ID    int
	Name  string
	Price int
}

type Order struct {
	ID       int
	FoodID   int
	FoodName string
	Quantity int
	Status   string
	CreatedAt string
}

var db *sql.DB
var templates *template.Template

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./food_ordering.db")
	if err != nil {
		log.Fatal(err)
	}

	foodsTableSQL := `CREATE TABLE IF NOT EXISTS foods (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, price INTEGER NOT NULL);`
	_, err = db.Exec(foodsTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	ordersTableSQL := `CREATE TABLE IF NOT EXISTS orders (id INTEGER PRIMARY KEY AUTOINCREMENT, food_id INTEGER NOT NULL, quantity INTEGER NOT NULL, status TEXT NOT NULL, create_at DATETIME DEFAULT CURRENT_TIMESTAMP);`
	_, err = db.Exec(ordersTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM foods").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		stmt, err := db.Prepare("INSERT INTO foods (name, price ) VALUES (?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		stmt.Exec("Nasi Goreng", 20000)
		stmt.Exec("Mie Ayam", 15000)
		stmt.Exec("Bakso", 18000)
		log.Println("Menu Awal Berhasil Ditambahkan")
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dashboard.html", nil)
}

func orderingHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, price FROM foods")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var foods []Food
	for rows.Next() {
		var food Food
		rows.Scan(&food.ID, &food.Name, &food.Price)
		foods = append(foods, food)
	}
	templates.ExecuteTemplate(w, "ordering.html", foods)
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT o.id, o.food_id, f.name, o.quantity, o.status, o.create_at
		FROM orders o
		JOIN foods f ON o.food_id = f.id
		ORDER BY o.create_at DESC`)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		rows.Scan(&order.ID, &order.FoodID, &order.FoodName, &order.Quantity, &order.Status, &order.CreatedAt)
		orders = append(orders, order)
	}
	templates.ExecuteTemplate(w, "orders.html", orders)
}

func placeOrderAPI(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FoodID int `json:"food_id"`
		Quantity int `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO orders (food_id, quantity, status) VALUES (?, ?, ?)",
		req.FoodID, req.Quantity, "pending")
	if err != nil {
		http.Error(w, "Failed to replace order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order place successfully!"})
}

func main() {
	initDB()
	defer db.Close()

	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()

	r.HandleFunc("/", dashboardHandler)
	r.HandleFunc("/pesan", orderingHandler)
	r.HandleFunc("/riwayat", ordersHandler)

	r.HandleFunc("/api/orders", placeOrderAPI).Methods("POST")

	fmt.Println("Server berjalan di http://localhos:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}