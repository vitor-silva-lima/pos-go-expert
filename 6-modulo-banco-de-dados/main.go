package main

import (
	"database/sql"

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
