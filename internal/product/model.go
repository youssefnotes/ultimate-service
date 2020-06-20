package product

import "time"

type Product struct {
	ProductId   string    `db:"product_id" json:"product_id"`
	Name        string    `db:"name" json:"name"`
	Price       int       `db:"price" json:"price"`
	Quantity    int       `db:"quantity" json:"quantity"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	DateUpdated time.Time `db:"date_updated" json:"date_updated"`
}
