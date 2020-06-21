package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/youssefnotes/ultimate-service/cmd/service/api/internal"
	"github.com/youssefnotes/ultimate-service/internal/platform"
	"github.com/youssefnotes/ultimate-service/internal/platform/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {

}
func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	log := log.New(os.Stdout, "sales service: ", log.LstdFlags)
	log.Println("run: started")
	defer log.Println("run: completed")

	// ====================================================================================
	// Configuration
	if os.Getenv("ENVIRONMENT") == "DEV" {
		viper.SetConfigName("config")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				log.Fatalln("init: ", err)
			} else {
				// Config file was found but another error was produced
			}
		}
	} else {
		viper.AutomaticEnv()
	}
	cfg := platform.Config{
		DB: platform.DBCfg{
			SSLmode:    viper.GetString("db_ssl_mode"),
			Timezone:   viper.GetString("db_time_zone"),
			Scheme:     viper.GetString("db_scheme"),
			Username:   viper.GetString("db_user_name"),
			Password:   viper.GetString("db_pass_word"),
			IP:         viper.GetString("db_ip"),
			Path:       viper.GetString("db_path"),
			DriverName: viper.GetString("db_driver_name"),
		},
		Web: platform.WebCfg{
			APIHost:         viper.GetString("app_ip"),
			Port:            viper.GetString("app_port"),
			ReadTimeout:     viper.GetDuration("read_timeout"),
			WriteTimeout:    viper.GetDuration("write_timeout"),
			ShutdownTimeout: viper.GetDuration("shutdown_timeout"),
		},
	}

	// ====================================================================================
	// Setting up dependency

	// ====================================================================================
	// Open Database
	db, err := database.Open(cfg.DB)
	if err != nil {
		return errors.Wrap(err, "run: opening db connection")
	}
	log.Println("run: db connected")
	defer db.Close()

	// ====================================================================================
	// Starting Server
	serverAddress := fmt.Sprintf("%s:%s", cfg.Web.APIHost, cfg.Web.Port)
	apiServer := http.Server{
		Addr:         serverAddress,
		Handler:      internal.API(log, db),
		TLSConfig:    nil,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		ErrorLog:     log,
	}
	serverError := make(chan error, 1)

	// Start service to listen for incoming requests
	go func() {
		log.Println("run: listening on: ", serverAddress)
		serverError <- apiServer.ListenAndServe()
	}()

	// Make channel and listen for shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Block and wait for shutdown
	select {
	case err := <-serverError:
		return errors.Wrap(err, "run: starting server")
	case <-shutdown:
		log.Println("run: starting shutdown")

	}

	// ====================================================================================
	// Shut Down Server

	return nil
}
