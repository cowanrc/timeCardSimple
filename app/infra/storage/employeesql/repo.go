package employeesql

import (
	"context"
	"database/sql"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
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
	return nil
}

func (r *Repo) RemoveEmployee(ctx context.Context, employeeID id.ID) error {
	return nil
}
