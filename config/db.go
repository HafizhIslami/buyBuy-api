package db

import (
	"database/sql"
	"fmt"

	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
)

// DB Configuration
type DB struct {
	// Address  string `env:"DB_ADDRESS,required"`
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	Username string `env:"DB_USERNAME,required"`
	Password string `env:"DB_PASSWORD,required"`
	DBName   string `env:"DB_NAME,required"`

	EnableLogging bool `env:"DB_ENABLE_LOGGING"`
}

// DB Connection checking
func GetConnection() error {
	d := &DB{}
	err := env.Parse(d)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", d.Username, d.Password, d.Host, d.Port, d.DBName))
	if err != nil {
		return err
	}

	defer db.Close()
	return nil
}