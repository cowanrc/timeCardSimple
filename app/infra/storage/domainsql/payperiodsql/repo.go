package payperiodsql

import (
	"context"
	"database/sql"
	"fmt"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/logger"
	"timeCardSimple/app/domain/payperiod"
	"timeCardSimple/app/infra/storage/domainsql/payperiodsql/queries"
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

func (r *Repo) GetPayPeriodByEmployeeID(ctx context.Context, employeeID id.ID) (*payperiod.PayPeriod, error) {
	query := queries.GetPayPeriodByEmployeeID

	row := r.sqlRepo.QueryRowContext(ctx, query, employeeID.GoString())

	options, err := scanPayPeriodOptions(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pay period for employeeID %v not found", employeeID)
		}
		return nil, err
	}

	return payperiod.NewWithOptions(options)
}

func (r *Repo) CreatePayPeriod(ctx context.Context, payPeriod *payperiod.PayPeriod) error {
	query := queries.CreatePayPeriod

	_, err := r.sqlRepo.ExecContext(ctx, query,
		payPeriod.ID().String(),
		payPeriod.EmployeeID().String(),
		payPeriod.StartDate(),
		payPeriod.DaysWorked(),
		payPeriod.TotalHours(),
	)
	if err != nil {
		logger.Error("error when trying to create employee pay period", err)
		return fmt.Errorf("could not create pay period: %v", err)
	}

	return nil
}

func scanPayPeriodOptions(rs *sql.Row) (options payperiod.Options, err error) {
	var startDate, endDate timesql.NullSQLTime

	err = rs.Scan(
		sqlrepo.ScanIntoID(&options.ID),
		sqlrepo.ScanIntoID(&options.EmployeeID),
		&options.StartDate,
		&options.EndDate,
		&options.DaysWorked,
		&options.TotalHours,
	)
	if err != nil {
		return options, err
	}

	options.StartDate = startDate.Domain()
	options.EndDate = endDate.Domain()

	return options, nil
}
