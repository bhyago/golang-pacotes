package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Product 1", 100.0)

	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	println("Product inserted:",
		product.ID, product.Name, product.Price)

	product.Price = 200.0
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	println("Product updated:",
		product.ID, product.Name, product.Price)

	product, err = selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product selected: %+v\n", product)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Printf("Product selected: %+v\n", product)
	}

	err = deleteOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	println("Product deleted:", product.ID)
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	product := &Product{}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []*Product{}

	for rows.Next() {
		product := &Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func deleteOneProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
