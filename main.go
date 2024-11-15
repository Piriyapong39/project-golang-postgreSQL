package main

import (
	"fmt"

	database "github.com/azujito/project-postgreSQL/config"
	Crude "github.com/azujito/project-postgreSQL/module"
)

type Product struct {
	Id          int
	Name        string
	Price       int
	Supplier_id int
}

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Create product
	product := Product{
		Name:        "TaoBao",
		Price:       25,
		Supplier_id: 2,
	}

	err = Crude.CreateProduct(db, product.Name, product.Price, product.Supplier_id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Create product successfully")

	// Get a single product
	rows, err := Crude.GetProduct(db, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)

	// Update product
	rows, err = Crude.UpdateProduct(db, 2, 1299)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)

	// Delete product
	err = Crude.DeleteProduct(db, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Delete product successfully")

	// Get all products
	results, err := Crude.GetProducts(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)

}
