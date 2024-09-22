package queries

const (
	GetEmployeeByID    = `SELECT * FROM employees WHERE employee_id = $1`
	GetEmployeeByEmail = `SELECT * FROM employees WHERE email = $1`
	InsertEmployee     = `INSERT INTO employees(id, first_name, last_name, email, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6);`
)
