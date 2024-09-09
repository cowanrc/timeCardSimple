package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	mysql_timecard_username = "mysql_timecard_username"
	mysql_timecard_password = "mysql_timecard_password"
	mysql_timecard_host     = "mysql_timecard_host"
	mysql_timecard_schema   = "mysql_timecard_schema"
	datbaseEnv              = "DATABASE_URL"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_timecard_username)
	password = os.Getenv(mysql_timecard_password)
	host     = os.Getenv(mysql_timecard_host)
	schema   = os.Getenv(mysql_timecard_schema)
)

func CreateDatabase() {
	ds := mustGetEnv(datbaseEnv)
	db := sqlx.MustConnect("postgres", ds)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("environment variable %s not set", key)
	}

	return value
}
