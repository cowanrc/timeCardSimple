package database

type Employee struct {
	Name        string `json:"name"`
	EmployeeID  int64  `json:"employeeID,omitempty"`
	ClockIn     string `json:"clockIn,omitempty"`
	ClockOut    string `json:"clockOut,omitempty"`
	TotalTime   string `json:"totalTime,omitempty"`
	DateOfBirth string `json:"dob"`
}

type TimeCard struct {
	Name       string `json:"name"`
	EmployeeID int64  `json:"employeeID,omitempty"`
	ClockIn    string `json:"clockIn,omitempty"`
	ClockOut   string `json:"clockOut,omitempty"`
	TotalTime  string `json:"totalTime,omitempty"`
}

type Employees []Employee
