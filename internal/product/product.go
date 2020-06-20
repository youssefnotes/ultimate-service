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
