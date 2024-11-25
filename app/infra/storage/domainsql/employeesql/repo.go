package employeesql

import (
	"context"
	"database/sql"
	"fmt"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/logger"
	"timeCardSimple/app/infra/storage/domainsql/employeesql/queries"
	"timeCardSimple/app/infra/storage/sqlrepo"

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

func (r *Repo) GetAllEmployees(ctx context.Context) ([]*employee.Employee, error) {
	query := queries.GetAllEmployees

	rows, err := r.sqlRepo.QueryContext(ctx, query)
	if err != nil {
		logger.Error("error getting all employees", err)
		return nil, fmt.Errorf("could not get all employees: %w", err)
	}
	defer rows.Close()

	var employees []*employee.Employee
	for rows.Next() {
		var options employee.Options
		err := rows.Scan(
			sqlrepo.ScanIntoID(&options.ID),
			&options.FirstName,
			&options.LastName,
			&options.Email,
			&options.CreatedAt,
			&options.UpdatedAt,
			// &options.PasswordHash,
		)
		if err != nil {
			logger.Error("error scanning employee row", err)
			return nil, fmt.Errorf("could not scan employee row: %w", err)
		}

		e, err := employee.NewWithOptions(options)
		if err != nil {
			return nil, fmt.Errorf("could not create employee instance: %w", err)
		}

		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		logger.Error("error during rows iteration", err)
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return employees, nil
}

func (r *Repo) GetEmployeeByID(ctx context.Context, employeeID id.ID) (*employee.Employee, error) {
	query := queries.GetEmployeeByID

	row := r.sqlRepo.QueryRowContext(ctx, query, employeeID.GoString())

	options, err := scanEmployeeOptions(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("employee with ID %v not found", employeeID)
		}
		return nil, err
	}

	return employee.NewWithOptions(options)
}

func (r *Repo) GetEmployeeByEmail(ctx context.Context, email string) (*employee.Employee, error) {
	return nil, nil
}
func (r *Repo) AddEmployee(ctx context.Context, employee *employee.Employee) error {
	query := queries.InsertEmployee

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
		return fmt.Errorf("could not save employee: %w", err)
	}

	return nil
}

func (r *Repo) RemoveEmployee(ctx context.Context, employeeID id.ID) error {
	query := queries.RemoveEmployee

	_, err := r.sqlRepo.ExecContext(ctx, query, employeeID.GoString())
	if err != nil {
		logger.Error("error removing employee", err)
		return fmt.Errorf("could not remove employee: %w", err)
	}
	return nil
}

func scanEmployeeOptions(rs *sql.Row) (options employee.Options, err error) {
	err = rs.Scan(
		sqlrepo.ScanIntoID(&options.ID),
		&options.FirstName,
		&options.LastName,
		&options.Email,
		&options.CreatedAt,
		&options.UpdatedAt,
	)

	return
}
