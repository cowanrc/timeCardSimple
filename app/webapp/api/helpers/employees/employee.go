package employees

import (
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/webapp/api"
)

func GetEmployeeId(employeeID string) (*id.ID, *api.RestErr) {
	empID, err := id.ParseString(employeeID)
	if err != nil {
		return nil, api.NewBadRequestError("Error parsing id from parameters")
	}

	return &empID, nil
}
