package weeklysummarysql

import (
	"context"
	"database/sql"
	"fmt"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/logger"
	"timeCardSimple/app/domain/weeklysummary"
	"timeCardSimple/app/infra/storage/domainsql/weeklysummarysql/queries"
	"timeCardSimple/app/infra/storage/sqlrepo"
	"timeCardSimple/app/infra/storage/timesql"
)

type Repo struct {
	sqlRepo *sql.DB
}

func New(sqlRepo *sql.DB) *Repo {
	return &Repo{
		sqlRepo: sqlRepo,
	}
}

func (r *Repo) GetWeeklySummaryByEmployeeID(ctx context.Context, employeeID id.ID) (*weeklysummary.WeeklySummary, error) {
	query := queries.GetWeeklySummaryByEmployeeID

	row := r.sqlRepo.QueryRowContext(ctx, query, employeeID.GoString())
	options, err := scanWeeklySummaryOptions(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("weekly summary for employeeID %v not found", employeeID)
		}
		return nil, err
	}
	return weeklysummary.NewWithOptions(options)
}

func (r *Repo) CreateWeeklySummary(ctx context.Context, weeklySummary *weeklysummary.WeeklySummary) error {
	query := queries.CreateWeeklySummary

	_, err := r.sqlRepo.ExecContext(ctx, query,
		weeklySummary.ID().String(),
		weeklySummary.EmployeeID().String(),
		weeklySummary.StartDate(),
		weeklySummary.DaysWorked(),
		weeklySummary.TotalHours(),
	)

	if err != nil {
		logger.Error("error when trying to create employee weekly summary", err)
		return fmt.Errorf("could not create weekly summary: %w", err)
	}

	return nil
}

func scanWeeklySummaryOptions(rs *sql.Row) (options weeklysummary.Options, err error) {
	var startDate timesql.NullSQLTime

	err = rs.Scan(
		sqlrepo.ScanIntoID(&options.ID),
		sqlrepo.ScanIntoID(&options.EmployeeID),
		&startDate,
		&options.DaysWorked,
		&options.TotalHours,
	)
	if err != nil {
		return options, err
	}

	options.StartDate = startDate.Domain()

	return options, nil
}
