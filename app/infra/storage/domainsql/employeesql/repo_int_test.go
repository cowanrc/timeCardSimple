package employeesql

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id/idtest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func setupMockDB(t *testing.T) (*Repo, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	repo := New(db)
	return repo, mock
}

func newTestEmployeeOptions() employee.Options {
	return employee.Options{
		ID:        idtest.MustNew(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestRepo_GetAllEmployees_Ok(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer mock.ExpectClose()

	options := newTestEmployeeOptions()
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "created_at", "updated_at"}).
		AddRow(options.ID.String(), options.FirstName, options.LastName, options.Email, options.CreatedAt, options.UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM employees").WillReturnRows(rows)

	ctx := context.Background()
	employees, err := repo.GetAllEmployees(ctx)
	require.NoError(t, err)
	require.Len(t, employees, 1)

	require.Equal(t, options.ID, employees[0].ID())
	require.Equal(t, options.FirstName, employees[0].FirstName())
	require.Equal(t, options.LastName, employees[0].LastName())
	require.Equal(t, options.Email, employees[0].Email())
}

func TestRepo_GetAllEmployees_Error(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer mock.ExpectClose()

	mock.ExpectQuery("SELECT (.+) FROM employees").WillReturnError(errors.New("query error"))

	ctx := context.Background()
	_, err := repo.GetAllEmployees(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "could not get all employees")
}

func TestRepo_GetEmployeeByID_Ok(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer mock.ExpectClose()

	options := newTestEmployeeOptions()
	row := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "created_at", "updated_at"}).
		AddRow(options.ID.String(), options.FirstName, options.LastName, options.Email, options.CreatedAt, options.UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM employees WHERE id = ?").
		WithArgs(options.ID.String()).
		WillReturnRows(row)

	ctx := context.Background()
	employee, err := repo.GetEmployeeByID(ctx, options.ID)
	require.NoError(t, err)
	require.Equal(t, options.ID, employee.ID())
}

func TestRepo_GetEmployeeByID_NotFound(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer mock.ExpectClose()

	nonExistentID := idtest.MustNew()
	mock.ExpectQuery("SELECT (.+) FROM employees WHERE id = ?").
		WithArgs(nonExistentID.GoString()).
		WillReturnError(sql.ErrNoRows)

	ctx := context.Background()
	_, err := repo.GetEmployeeByID(ctx, nonExistentID)
	require.Error(t, err)
	if err != sql.ErrNoRows {
		t.Errorf("incorrect error, received %v, expected %v", err, sql.ErrNoRows)
	}
}

func TestRepo_AddEmployee_Ok(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer mock.ExpectClose()

	e := newTestEmployeeOptions()
	mock.ExpectExec("INSERT INTO employees").
		WithArgs(e.ID.String(), e.FirstName, e.LastName, e.Email, e.CreatedAt, e.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	newEmployee, err := employee.NewWithOptions(e)
	require.NoError(t, err)
	err = repo.AddEmployee(ctx, newEmployee)
	require.NoError(t, err)
}

func TestRepo_RemoveEmployee_Ok(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer mock.ExpectClose()

	employeeID := idtest.MustNew()
	mock.ExpectExec("DELETE FROM employees WHERE id = ?").
		WithArgs(employeeID.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	err := repo.RemoveEmployee(ctx, employeeID)
	require.NoError(t, err)
}

func TestRepo_RemoveEmployee_NotFound(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer mock.ExpectClose()

	employeeID := idtest.MustNew()
	mock.ExpectExec("DELETE FROM employees WHERE id = ?").
		WithArgs(employeeID.String()).
		WillReturnError(errors.New("no rows affected"))

	ctx := context.Background()
	err := repo.RemoveEmployee(ctx, employeeID)
	require.Error(t, err)
	require.Contains(t, err.Error(), "could not remove employee")
}
