package employees

import (
	"context"
	"database/sql"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
)

type Repo struct {
	Client *sql.DB
}

func (r *Repo) GetEmployeeByID(ctx context.Context, employeeID id.ID) (*employee.Employee, error)
