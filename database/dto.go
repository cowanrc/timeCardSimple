package database

type Employee struct {
	Name        string `json:"name,omitempty"`
	EmployeeID  int64  `json:"employeeID,omitempty"`
	ClockIn     string `json:"clockIn,omitempty"`
	ClockOut    string `json:"clockOut,omitempty"`
	TotalTime   string `json:"totalTime,omitempty"`
	DateOfBirth string `json:"dob,omitempty"`
}
