package postgres

import (
	"acme/cmd/config"
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg config.Config) (*gorm.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%s", cfg.DbHost, cfg.DbPort),
		User:   url.UserPassword(cfg.DbUser, cfg.DbPassword),
		Path:   cfg.DbName,
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
