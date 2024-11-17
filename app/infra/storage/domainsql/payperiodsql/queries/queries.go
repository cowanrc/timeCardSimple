package queries

const (
	// GET
	GetPayPeriodByEmployeeID = `SELECT * FROM pay_periods WHERE employee_id=$1;`

	// ADD
	CreatePayPeriod = `INSERT INTO pay_periods(id, employee_id, period_start_date, total_days_worked, total_hours)
	VALUES ($1, $2, $3, $4, $5);`
)
