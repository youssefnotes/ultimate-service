package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //register postgres driver
	"net/url"
)

type Config struct {
	Sslmode        string
	Timezone       string
	DB_scheme      string
	DB_user_name   string
	DB_pass_word   string
	DB_ip          string
	DB_path        string
	DB_driver_name string
}

func Open(cfg Config) (*sqlx.DB, error) {
	q := url.Values{}
	q.Set("sslmode", cfg.Sslmode)
	q.Set("timezone", cfg.Timezone)
	u := url.URL{
		Scheme:   cfg.DB_scheme,
		User:     url.UserPassword(cfg.DB_user_name, cfg.DB_pass_word),
		Host:     cfg.DB_ip,
		Path:     cfg.DB_path,
		RawQuery: q.Encode(),
	}
	return sqlx.Open(cfg.DB_driver_name, u.String())
}
