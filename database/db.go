package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_timecard_username = "mysql_timecard_username"
	mysql_timecard_password = "mysql_timecard_password"
	mysql_timecard_host     = "mysql_timecard_host"
	mysql_timecard_schema   = "mysql_timecard_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_timecard_username)
	password = os.Getenv(mysql_timecard_password)
	host     = os.Getenv(mysql_timecard_host)
	schema   = os.Getenv(mysql_timecard_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	log.Println(username, password, host, schema)

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
