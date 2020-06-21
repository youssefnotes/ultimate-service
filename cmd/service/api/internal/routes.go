package internal

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func API(log *log.Logger, db *sqlx.DB) http.Handler {
	p := ProductService{Log: log, DB: db}

	return http.HandlerFunc(p.List)
}
