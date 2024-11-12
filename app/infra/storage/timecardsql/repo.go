package timecardsql

import (
	"context"
	"database/sql"
	"fmt"
	"timeCardSimple/app/domain/logger"
	"timeCardSimple/app/domain/timecard"
	"timeCardSimple/app/infra/storage/timecardsql/queries"
)

type Repo struct {
	sqlRepo *sql.DB
}

func New(sqlRepo *sql.DB) *Repo {
	return &Repo{
		sqlRepo: sqlRepo,
	}
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
