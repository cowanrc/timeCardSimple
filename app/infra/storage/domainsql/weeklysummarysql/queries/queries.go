package queries

const (
	// GET
	GetWeeklySummaryByEmployeeID = `SELECT * FROM weekly_summary WHERE employee_id=$1`

	// ADD
	CreateWeeklySummary = `INSERT INTO weekly_summary (id, employee_id, week_start_date, days_worked, total_hours)
	VALUES ($1, $2, $3, $4, $5);`

	// SET
)
