package main

import "fmt"

type User struct{
	FirstName string
	LastName string
	Age int
}

func (u User) FullName() string{
	return u.FirstName + " " + u.LastName
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func main() {
	user := User{"Budi", "santosol", 20}
	fmt.Println("Nama Lengkap Saya:", user.FullName())
	fmt.Println("Apakah Sudah dewasa:", user.IsAdult())
}