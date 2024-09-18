package queries

const (
	GetEmployeeByID    = `SELECT * FROM employees WHERE employee_id = $1`
	GetEmployeeByEmail = `SELECT * FROM employees WHERE email = $1`
)
