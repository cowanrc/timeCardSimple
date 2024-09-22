package employeesql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/logger"
	"timeCardSimple/app/infra/storage/employeesql/queries"

	_ "github.com/lib/pq"
)

type Repo struct {
	sqlRepo *sql.DB
}

// New returns a new Repo.
func New(sqlRepo *sql.DB) *Repo {
	return &Repo{
		sqlRepo: sqlRepo,
	}
}

func (r *Repo) GetEmployeeByID(ctx context.Context, employeeID id.ID) (*employee.Employee, error) {
	return nil, nil
}

func (r *Repo) GetEmployeeByEmail(ctx context.Context, email string) (*employee.Employee, error) {
	return nil, nil
}
func (r *Repo) AddEmployee(ctx context.Context, employee *employee.Employee) error {
	// Use context-aware queries and reusing prepared statements
	query := queries.InsertEmployee

	// Use named parameters for better readability if supported
	_, err := r.sqlRepo.ExecContext(ctx, query,
		employee.ID().String(),
		employee.FirstName(),
		employee.LastName(),
		employee.Email(),
		employee.CreatedAt(),
		employee.UpdatedAt(),
	)

	if err != nil {
		logger.Error("error when trying to save employee", err)
		return fmt.Errorf("could not save employee: %w", err) // Wrapping error for better tracing
	}

	log.Println("SUCCESSFUL AT DB")
	return nil
}

func (r *Repo) RemoveEmployee(ctx context.Context, employeeID id.ID) error {
	return nil
}
