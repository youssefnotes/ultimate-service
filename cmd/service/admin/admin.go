package main

import (
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/youssefnotes/ultimate-service/internal/platform/database"
	"github.com/youssefnotes/ultimate-service/internal/schema"
	"log"
)

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	// ============================================================================================================
	// setup dependency

	// open database
	db, err := database.Open()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connected")
	defer db.Close()

	flag.Parse()
	switch flag.Arg(0) {
	case "migrate":
		//migrations
		if err = schema.Migrate(db); err != nil {
			log.Fatalln(err)
		}
		log.Println("Migration complete")
		return
	case "seed":
		if err = schema.Seed(db); err != nil {
			log.Fatalln(err)
		}
		log.Println("Seed complete")
		return
	}
}
