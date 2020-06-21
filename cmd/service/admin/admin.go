package main

import (
	"flag"
	"github.com/spf13/viper"
	"github.com/youssefnotes/ultimate-service/internal/platform/database"
	"github.com/youssefnotes/ultimate-service/internal/schema"
	"log"
	"os"
)

func init() {
	if os.Getenv("ENVIRONMENT") == "DEV" {
		viper.SetConfigName("config")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Panic(err)
			} else {
				// Config file was found but another error was produced
			}
		}
	} else {
		viper.AutomaticEnv()
	}
}
func main() {
	// ============================================================================================================
	// setup dependency

	// open database
	db, err := database.Open(database.Config{
		Sslmode:        viper.GetString("db_ssl_mode"),
		Timezone:       viper.GetString("db_time_zone"),
		DB_scheme:      viper.GetString("db_scheme"),
		DB_user_name:   viper.GetString("db_user_name"),
		DB_pass_word:   viper.GetString("db_pass_word"),
		DB_ip:          viper.GetString("db_ip"),
		DB_path:        viper.GetString("db_path"),
		DB_driver_name: viper.GetString("db_driver_name"),
	})
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
