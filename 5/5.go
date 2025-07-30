package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtkey = []byte("inikuncirahasianya")

type User struct {
	Username string `json: "username"`
	Password string `json: "password`
}

type Claims struct {
	Username string `json: "username"`
	jwt.StandardClaims
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "/.users.db")
	if err != nil {
		log.Fatal("Gagal Membuka Database: %v", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Gagal Membuat Table: %v", err)
	}
	log.Println("Database dan Table 'users' berhasil diinisialisasi atau sudah ada")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body) .Decode(&user); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		http.Error(w, "Database ERROR", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, string(hashedPassword))
	if err != nil {
		http.Error(w, "Username alreadu exist or other database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"massage": "User Registered Successfully"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Credential", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error: Failed to retrieve user", http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invaid Credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := generateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generate Token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	return tokenString, err
}

func validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing Authorization Token", http.StatusUnauthorized)
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtkey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid Token Signature", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selamat Datang di Area yang Dilundungi! Anda Berhasil Mengaksesnya dengan JWT yang Valid."))
}

func main() {
	initDB()

	r := mux.NewRouter()

	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	r.Handle("/protected", validateToken(http.HandlerFunc(protectedHandler))).Methods("GET")

	fmt.Println("Server Berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}