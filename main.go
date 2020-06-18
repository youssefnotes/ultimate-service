package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/youssefnotes/ultimate-service/schema"
	"log"
	"net/http"
	"net/url"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	log.Println("main: Started")
	defer log.Println("main: Completed")

	// ============================================================================================================
	// setup dependency

	// open database
	db, err := openDB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("db connected")
	defer db.Close()

	flag.Parse()
	switch flag.Arg(0) {
	case "migrate":
		//migrations
		if err = schema.Migrate(db); err != nil {
			log.Fatalln(err)
		}
		log.Println("db migration complete")
		return
	case "seed":
		if err = schema.Seed(db); err != nil {
			log.Fatalln(err)
		}
		log.Println("db seed complete")
		return
	}

	address := fmt.Sprintf("%s:%s", os.Getenv("app_ip"), os.Getenv("app_port"))
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

func openDB() (*sqlx.DB, error) {
	q := url.Values{}
	q.Set("sslmode", os.Getenv("db_ssl_mode"))
	q.Set("timezone", os.Getenv("db_time_zone"))
	u := url.URL{
		Scheme:   os.Getenv("db_scheme"),
		User:     url.UserPassword(os.Getenv("db_user_name"), os.Getenv("db_pass_word")),
		Host:     os.Getenv("db_ip"),
		Path:     os.Getenv("db_path"),
		RawQuery: q.Encode(),
	}

	return sqlx.Open(os.Getenv("db_driver_name"), u.String())
}
