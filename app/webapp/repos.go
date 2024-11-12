package webapp

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/timecard"
	"timeCardSimple/app/infra/storage/employeesql"
	"timeCardSimple/app/infra/storage/timecardsql"

	_ "github.com/lib/pq"
)

const (
	datbaseEnv        = "DATABASE_URL"
	sqlOpenDriverName = "postgres"
)

type Repos struct {
	Employee employee.Repo
	Timecard timecard.Repo
}

func BuildRepos() (*sql.DB, *Repos, error) {
	db, err := buildDb()
	if err != nil {
		panic("Couldn't start database")
	}

	repos := &Repos{
		Employee: employeesql.New(db),
		Timecard: timecardsql.New(db),
	}

	return db, repos, nil
}

func buildDb() (*sql.DB, error) {
	dbConn := os.Getenv(datbaseEnv)
	u, err := url.Parse(dbConn)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", u.String())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

	return db, err
}
