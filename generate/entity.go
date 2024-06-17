package main

//go:generate go run generate_queries.go templete.go
type User struct {
	ID   int
	Name string
}
type Product struct {
	ID    int
	Name  string
	Price float64
}
