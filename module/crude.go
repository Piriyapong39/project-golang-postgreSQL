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

type ProductAndSupplier struct {
	Name          string
	Price         int
	Supplier_Name string
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

func GetProducts(db *sql.DB) ([]Product, error) {
	var products []Product
	rows, err := db.Query("SELECT * FROM tb_products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := new(Product)
		rows.Scan(&p.Id, &p.Name, &p.Price, &p.Supplier_id)
		products = append(products, *p)
	}
	return products, nil
}

func GetProductsAndSupplierName(db *sql.DB) ([]ProductAndSupplier, error) {
	var products []ProductAndSupplier

	rows, err := db.Query(`
		SELECT 
			p.name,
			p.price,
			s.name
		FROM tb_products p
		INNER JOIN 
			tb_supplier s ON p.supplier_id = s.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := new(ProductAndSupplier)
		err := rows.Scan(&p.Name, &p.Price, &p.Supplier_Name)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}

// func GetProductsAndSupplierName(db *sql.DB) ([]ProductAndSupplier, error) {
// 	var products []ProductAndSupplier
// 	rows, err := db.Query(`
//         SELECT
//             p.name,
//             p.price,
//             s.name
//         FROM
//             tb_products p
//         INNER JOIN
//             tb_supplier s ON p.supplier_id = s.id
//     `)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var p ProductAndSupplier
// 		err := rows.Scan(&p.ProductName, &p.ProductPrice, &p.SupplierName)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, p)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return products, nil
// }
