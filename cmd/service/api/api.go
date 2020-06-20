package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/youssefnotes/ultimate-service/cmd/service/api/internal"
	"github.com/youssefnotes/ultimate-service/internal/platform/database"
	"log"
	"net/http"
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
	db, err := database.Open()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("db connected")
	defer db.Close()

	address := fmt.Sprintf("%s:%s", os.Getenv("app_ip"), os.Getenv("app_port"))
	if err := http.ListenAndServe(address, http.HandlerFunc((&internal.ProductService{DB: db}).List)); err != nil {
		log.Fatalln(err)
	}
}
