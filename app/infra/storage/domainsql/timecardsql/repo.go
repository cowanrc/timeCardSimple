package timecardsql

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/logger"
	"timeCardSimple/app/domain/timecard"
	"timeCardSimple/app/infra/storage/domainsql/timecardsql/queries"
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

func (r *Repo) GetTimecardByEmployeeID(ctx context.Context, employeeID id.ID) (*timecard.Timecard, error) {
	query := queries.GetTimecardByEmployeeID

	row := r.sqlRepo.QueryRowContext(ctx, query, employeeID.GoString())

	options, err := scanTimecardOptions(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("timecard for employeeID %v not found", employeeID)
		}
		return nil, err
	}

	return timecard.NewWithOptions(options)
}

func (r *Repo) CreateEmployeeTimecard(ctx context.Context, timecard *timecard.Timecard) error {
	query := queries.CreateTimecard

	_, err := r.sqlRepo.ExecContext(ctx, query,
		timecard.ID().String(),
		timecard.EmployeeID().String(),
		timecard.CreatedAt(),
		timecard.UpdatedAt(),
	)

	if err != nil {
		logger.Error("error when trying to create employee timecard", err)
		return fmt.Errorf("could not create timecard: %w", err)
	}

	return nil
}

func (r *Repo) ClockInEmployee(
	ctx context.Context,
	employeeID id.ID,
	startTime time.Time,
	weekStartDate *time.Time,
	BiWeekStartDate *time.Time,
) error {
	query := queries.ClockInEmployee

	_, err := r.sqlRepo.ExecContext(ctx, query,
		startTime,
		weekStartDate,
		BiWeekStartDate,
		employeeID.String(),
	)

	if err != nil {
		logger.Error("error when trying to clock in employee", err)
		return fmt.Errorf("could not clock in employee: %w", err)
	}

	return nil
}

func scanTimecardOptions(rs *sql.Row) (options timecard.Options, err error) {
	var startTime, endTime, weekStartDate, biWeeklyPeriodStart timesql.NullSQLTime
	var duration sql.NullFloat64

	err = rs.Scan(
		sqlrepo.ScanIntoID(&options.ID),
		sqlrepo.ScanIntoID(&options.EmployeeID),
		&startTime,
		&endTime,
		&options.CreatedAt,
		&options.UpdatedAt,
		&duration,
		&weekStartDate,
		&biWeeklyPeriodStart,
	)
	if err != nil {
		return options, err
	}

	options.StartTime = startTime.Domain()
	options.EndTime = endTime.Domain()
	options.WeekStartDate = weekStartDate.Domain()
	options.BiWeeklyPeriodStart = biWeeklyPeriodStart.Domain()

	if duration.Valid {
		options.Duration = &duration.Float64
	} else {
		options.Duration = nil
	}

	return options, nil
}
