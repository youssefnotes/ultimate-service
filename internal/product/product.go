package product

import (
	"github.com/jmoiron/sqlx"
)

func List(db *sqlx.DB) ([]Product, error) {
	const q = `SELECT * FROM products`
	list := []Product{}
	if err := db.Select(&list, q); err != nil {
		return nil, err
	}
	return list, nil
}

func Retrieve(db *sqlx.DB, id string) (*Product, error) {
	var p Product
	q := `SELECT product_id, name, price, quantity, date_created, date_updated FROM products WHERE product_id = $1`
	err := db.Get(&p, q, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
