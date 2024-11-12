package queries

const (
	// ADD
	CreateTimecard = `INSERT INTO timecard (id, employee_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4);`
)
