package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //register postgres driver
	"net/url"
	"os"
)

func Open() (*sqlx.DB, error) {
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
