package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	address := "localhost:8000"
	if err := http.ListenAndServe(address, http.HandlerFunc(getProducts)); err != nil {
		log.Fatalln(err)
	}
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}

func getProducts(writer http.ResponseWriter, request *http.Request) {

	products := []Product{}
	if true {
		products = append(products, Product{Name: "Comic Books", Price: 30, Qty: 20})
		products = append(products, Product{Name: "Medical Books", Price: 100, Qty: 10})
	}
	resp, err := json.Marshal(products)
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
