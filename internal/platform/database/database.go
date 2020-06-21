package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //register postgres driver
	"github.com/youssefnotes/ultimate-service/internal/platform"
	"net/url"
)

func Open(dbCfg platform.DBCfg) (*sqlx.DB, error) {
	q := url.Values{}
	q.Set("sslmode", dbCfg.SSLmode)
	q.Set("timezone", dbCfg.Timezone)
	u := url.URL{
		Scheme:   dbCfg.Scheme,
		User:     url.UserPassword(dbCfg.Username, dbCfg.Password),
		Host:     dbCfg.IP,
		Path:     dbCfg.Path,
		RawQuery: q.Encode(),
	}
	return sqlx.Open(dbCfg.DriverName, u.String())
}
