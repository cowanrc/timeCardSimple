package queries

const (
	// GET
	GetEmployeeByID    = `SELECT * FROM employees WHERE id = $1`
	GetEmployeeByEmail = `SELECT * FROM employees WHERE email = $1`
	GetAllEmployees    = `SELECT * FROM employees`

	// ADD
	InsertEmployee = `INSERT INTO employees(id, first_name, last_name, email, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6);`

	// REMOVE
	RemoveEmployee = `DELETE FROM employees WHERE id = $1`
)
