package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/youssefnotes/ultimate-service/cmd/service/api/internal"
	"github.com/youssefnotes/ultimate-service/internal/platform/database"
	"log"
	"net/http"
	"os"
)

func init() {
	if os.Getenv("ENVIRONMENT") == "DEV" {
		viper.SetConfigName("config")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
			} else {
				// Config file was found but another error was produced
			}
		}
	} else {
		viper.AutomaticEnv()
	}
}
func main() {
	log.Println("main: Started")
	defer log.Println("main: Completed")

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
	log.Println("db connected")
	defer db.Close()

	address := fmt.Sprintf("%s:%s", viper.GetString("app_ip"), viper.GetString("app_port"))
	if err := http.ListenAndServe(address, http.HandlerFunc((&internal.ProductService{DB: db}).List)); err != nil {
		log.Fatalln(err)
	}
}
