package db

import (
	"database/sql"
	"fmt"
	"net/url"
)

func Initlize(cfgFile string) (*sql.DB, error) {
	fmt.Println("Initializing database...")

	config, err := loadConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(config.SourceDB.User, config.SourceDB.Password),
		Host:     fmt.Sprintf("%s:%d", config.SourceDB.Host, config.SourceDB.Port),
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		return nil, err
	}
	return db, nil
}
