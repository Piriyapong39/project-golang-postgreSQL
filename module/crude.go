package Crude

import (
	"database/sql"
)

type Product struct {
	Id          int
	Name        string
	Price       int
	Supplier_id int
}

func CreateProduct(db *sql.DB, name string, price int, supplier_id int) error {
	_, err := db.Exec("INSERT INTO tb_products (name, price, supplier_id) VALUES ($1, $2, $3)",
		name, price, supplier_id)
	if err != nil {
		return err
	}
	return nil
}

func GetProduct(db *sql.DB, id int) (Product, error) {
	productTarget := new(Product)
	err := db.QueryRow("SELECT * FROM tb_products WHERE id=$1", id).
		Scan(&productTarget.Id, &productTarget.Name, &productTarget.Price, &productTarget.Supplier_id)
	if err != nil {
		return Product{}, err
	}
	return *productTarget, nil
}

func UpdateProduct(db *sql.DB, id int, price int) (Product, error) {
	product := new(Product)
	err := db.QueryRow("UPDATE tb_products SET price=$1 WHERE id=$2 RETURNING id, name, price, supplier_id", price, id).
		Scan(&product.Id, &product.Name, &product.Price, &product.Supplier_id)
	if err != nil {
		return Product{}, err
	}
	return *product, nil
}

func DeleteProduct(db *sql.DB, id int) error {
	_, err := db.Query("DELETE FROM tb_products WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}