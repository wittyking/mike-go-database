package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)

	// fmt.Println(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	fmt.Println("Connection Database Successful")
	////create
	// err = createProduct(&Product{Name: "Go product2", Price: 444})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	////Read
	// product, err := getProduct(2)
	// fmt.Println("Get Successful !", product)
	////update
	// product, err := updateProduct(2, &Product{Name: "Lopster111", Price: 400})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	////Delete
	// err = deleteProduct(4)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	products, err := getProducts()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get All Products !", products)

	// print("Create Successful!")
}
