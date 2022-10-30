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

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// product := NewProduct("macbook PRO", 17779.99)
	// err = InsertProduct(db, product)
	// if err != nil {
	// 	panic(err)
	// }
	// product.Name = "Macbook Air M1"
	// product.ID = "14ba438d-f200-4a1b-b611-8e0b4deb0019"
	// UpdateProduct(db, product)

	// p, err := SelectProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Produto: %v, Price: R$%.2f\n", p.Name, p.Price)

	err = DeleteProduct(db, "64c8fe50-18fd-40e7-950d-df8a651a4271")
	if err != nil {
		panic(err)
	}

	products, err := SelectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, v := range products {
		fmt.Printf("Produto: %v, Price: R$%.2f\n", v.Name, v.Price)
	}

}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func InsertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?,?,?)")
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

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
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

func SelectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func SelectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func DeleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
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
