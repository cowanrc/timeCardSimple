package queries

const (
	// GET
	GetTimecardByEmployeeID = `Select * FROM timecard WHERE employee_id = $1;`

	// ADD
	CreateTimecard = `INSERT INTO timecard (id, employee_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4);`

	// SET
	ClockInEmployee = `UPDATE timecard SET start_time = $1, week_start_date = $2, bi_weekly_period_start = $3
	WHERE employee_id = $4;`
)
