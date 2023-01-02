package api

import (
	"fmt"
	"strconv"
	"time"
	"timeCardSimple/errors"
)

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetEmployeeId(employeeIdParam string) (int64, *errors.RestErr) {
	employeeId, empErr := strconv.ParseInt(employeeIdParam, 10, 64)

	if empErr != nil {
		return 0, errors.NewBadRequestError("user id should be an integer")
	}

	return employeeId, nil
}

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}

func CalculateTotalTime(clockIn string, clockOut string) string {
	clockInTime, _ := time.Parse("Mon Jan _2 15:04:05 MST 2006", clockIn)
	clockOutTime, _ := time.Parse("Mon Jan _2 15:04:05 MST 2006", clockOut)
	totalTime := RoundDuration(clockOutTime.Sub(clockInTime))

	return totalTime
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
