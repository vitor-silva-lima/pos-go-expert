package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Product struct {
	ProductId uuid.UUID
	Name      string
	Price     float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ProductId: uuid.New(),
		Name:      name,
		Price:     price,
	}
}

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Product 2", 470.50)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Name = "Product 2 Updated"
	product.Price = 500.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	product, err = selectProduct(db, product.ProductId.String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("One Product: %+v\n", product)
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Printf("Product: %+v\n", product)
	}
	err = deleteProduct(db, product.ProductId.String())
	if err != nil {
		panic(err)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (product_id, name, price) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ProductId, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = $1, price = $2 WHERE product_id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT product_id, name, price FROM products WHERE product_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	product := &Product{}
	err = row.Scan(&product.ProductId, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("SELECT product_id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []*Product{}
	for rows.Next() {
		product := &Product{}
		err = rows.Scan(&product.ProductId, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE product_id = $1")
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
