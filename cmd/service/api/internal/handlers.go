package internal

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/youssefnotes/ultimate-service/internal/product"
	"log"
	"net/http"
)

type ProductService struct {
	DB *sqlx.DB
}

func (p ProductService) List(writer http.ResponseWriter, request *http.Request) {

	list, err := product.List(p.DB)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("product list ", err)
		return
	}

	resp, err := json.Marshal(list)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("marshal products ", err)
		return
	}

	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(resp)
	if err != nil {
		log.Println("get products ", err)
		return
	}
	return
}
