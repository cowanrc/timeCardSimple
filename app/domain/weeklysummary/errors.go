package weeklysummary

import "fmt"

const packagePrefix = "weeklysummary"

var (
	ErrUnsetID         = fmt.Errorf("%s: id is empty", packagePrefix)
	ErrUnsetEmployeeID = fmt.Errorf("%s: employeeID is empty", packagePrefix)
)
