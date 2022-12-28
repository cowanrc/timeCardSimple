package employee

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"timeCardSimple/errors"
)

func getEmployeeId(employeeIdParam string) (int64, *errors.RestErr) {
	employeeId, empErr := strconv.ParseInt(employeeIdParam, 10, 64)

	if empErr != nil {
		return 0, errors.NewBadRequestError("user id should be an integer")
	}

	return employeeId, nil
}

func employeeClockIn(id int) {
	employee := TimeCard[id]
	employee.ClockIn = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")
	log.Printf("Employee clocked in at: %s", employee.ClockIn)
}

func employeeClockOut(id int) {
	employee := TimeCard[id]
	employee.ClockOut = time.Now().UTC().Format("Mon Jan _2 15:04:05 MST 2006")
	log.Printf("Employee clocked out at: %s", employee.ClockOut)
}

func employeeExists(ID int) bool {
	_, ok := TimeCard[ID]
	return ok
}

func employeeTotalTime(clockIn string, clockOut string) string {
	clockInTime, _ := time.Parse("Mon Jan _2 15:04:05 MST 2006", clockIn)
	clockOutTime, _ := time.Parse("Mon Jan _2 15:04:05 MST 2006", clockOut)
	totalTime := roundDuration(clockOutTime.Sub(clockInTime))
	return totalTime
}

func roundDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	s := d / time.Second

	totalTime := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	return totalTime
}
