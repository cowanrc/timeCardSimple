package api

import (
	"fmt"
	"strconv"
	"time"
	"timeCardSimple/errors"
)

func GetEmployeeId(employeeIdParam string) (int64, *errors.RestErr) {
	employeeId, empErr := strconv.ParseInt(employeeIdParam, 10, 64)

	if empErr != nil {
		return 0, errors.NewBadRequestError("user id should be an integer")
	}

	return employeeId, nil
}

func RoundDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	s := d / time.Second

	totalTime := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	return totalTime
}
