package webapp

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/payperiod"
	"timeCardSimple/app/domain/timecard"
	"timeCardSimple/app/domain/weeklysummary"
	"timeCardSimple/app/infra/storage/domainsql/employeesql"
	"timeCardSimple/app/infra/storage/domainsql/payperiodsql"
	"timeCardSimple/app/infra/storage/domainsql/timecardsql"
	"timeCardSimple/app/infra/storage/domainsql/weeklysummarysql"

	_ "github.com/lib/pq"
)

const (
	datbaseEnv        = "DATABASE_URL"
	sqlOpenDriverName = "postgres"
)

type Repos struct {
	Employee      employee.Repo
	Timecard      timecard.Repo
	WeeklySummary weeklysummary.Repo
	PayPeriod     payperiod.Repo
}

func BuildRepos() (*sql.DB, *Repos, error) {
	db, err := buildDb()
	if err != nil {
		panic("Couldn't start database")
	}

	repos := &Repos{
		Employee:      employeesql.New(db),
		Timecard:      timecardsql.New(db),
		WeeklySummary: weeklysummarysql.New(db),
		PayPeriod:     payperiodsql.New(db),
	}

	return db, repos, nil
}

func buildDb() (*sql.DB, error) {
	dbConn := os.Getenv(datbaseEnv)
	u, err := url.Parse(dbConn)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(sqlOpenDriverName, u.String())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

	return db, err
}
