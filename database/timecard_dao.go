package database

import (
	"strings"
	"timeCardSimple/errors"
	"timeCardSimple/logger"
)

const (
	queryGetEmployeeSimple   = "SELECT employeeID, name FROM timecard WHERE employeeID=?;"
	queryClockIn             = "UPDATE timecard SET clockIn=? WHERE employeeID=?;"
	queryGetClockIn          = "SELECT employeeID, name, clockIn FROM timecard WHERE employeeID=?;"
	queryClockOut            = "UPDATE timecard SET clockOut=? WHERE employeeID=?;"
	queryGetClockInClockOut  = "Select employeeID, name, clockIn, clockOut FROM timecard WHERE employeeID=?"
	queryTotalTime           = "UPDATE timecard SET totalTime=? WHERE employeeID=?;"
	queryGetEmployeeTimecard = "SELECT employeeID, name, clockIn, clockOut, totalTime FROM timecard WHERE employeeID=?;"
)

func GetSimple(employee *TimeCard) *errors.RestErr {
	stmt, err := Client.Prepare(queryGetEmployeeSimple)
	if err != nil {
		logger.Error("error when trying to prepare get employee statement", err)
		return errors.NewInternalServerError("error trying to get employee")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		return errors.NewNotFoundError("employee might not exist in the system")
	}

	return nil
}
func EmployeeClockIn(employee *TimeCard) *errors.RestErr {
	stmt, err := Client.Prepare(queryClockIn)
	if err != nil {
		logger.Error("error when trying to prepare clockIn employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(employee.ClockIn, employee.EmployeeID)
	if err != nil {
		logger.Error("error when trying to clockIn employee", err)
		return errors.NewInternalServerError("error has occured when trying to clock in employee")
	}

	return nil
}

func EmployeeClockOut(employee *TimeCard) *errors.RestErr {
	stmt, err := Client.Prepare(queryClockOut)
	if err != nil {
		logger.Error("error when trying to prepare clockOut employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(employee.ClockOut, employee.EmployeeID)
	if err != nil {
		logger.Error("error when trying to clockOut employee", err)
		return errors.NewInternalServerError("error has occured when trying to clock out employee")
	}

	return nil
}

func EmployeeTotalTime(employee *TimeCard) *errors.RestErr {
	stmt, err := Client.Prepare(queryTotalTime)
	if err != nil {
		logger.Error("error when trying to prepare totalTime employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(employee.TotalTime, employee.EmployeeID)
	if err != nil {
		logger.Error("error when trying to insert total time of employee", err)
		return errors.NewInternalServerError("error when trying to calculate total time for employee")
	}

	return nil
}

func GetClockIn(employee *TimeCard) *errors.RestErr {
	stmt, err := Client.Prepare(queryGetClockIn)
	if err != nil {
		logger.Error("error when trying to prepare get employee statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name, &employee.ClockIn); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return errors.NewNotFoundError("employee might not exist in the system")
		} else {
			return errors.NewBadRequestError("employee must clock in before clocking out")
		}
	}

	return nil
}

func GetClockInClockOut(employee *TimeCard) *errors.RestErr {
	stmt, err := Client.Prepare(queryGetClockInClockOut)
	if err != nil {
		logger.Error("error when trying to prepare get employee clock in/clock out statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name, &employee.ClockIn, &employee.ClockOut); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return errors.NewNotFoundError("employee might not exist in the system")
		} else {
			return errors.NewBadRequestError("make sure employee has clocked in and has clocked out before requesting total time")
		}
	}

	return nil
}

func GetTime(employee *TimeCard) *errors.RestErr {
	stmt, err := Client.Prepare(queryGetEmployeeTimecard)
	if err != nil {
		logger.Error("error when trying to prepare get employee time card statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name, &employee.ClockIn, &employee.ClockOut, &employee.TotalTime); getErr != nil {
		logger.Error("error when trying to get employee time card by ID", getErr)
		return errors.NewNotFoundError("employee might not exist in the system")
	}

	return nil
}
