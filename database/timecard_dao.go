package database

import (
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

func (employee *TimeCard) Get() *errors.RestErr {
	stmt, err := Client.Prepare(queryGetEmployeeSimple)
	if err != nil {
		logger.Error("error when trying to prepare get employee statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
func (employee *TimeCard) EmployeeClockIn() *errors.RestErr {
	stmt, err := Client.Prepare(queryClockIn)
	if err != nil {
		logger.Error("error when trying to prepare clockIn employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(employee.ClockIn, employee.EmployeeID)
	if err != nil {
		logger.Error("error when trying to clockIn employee", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (employee *TimeCard) EmployeeClockOut() *errors.RestErr {
	stmt, err := Client.Prepare(queryClockOut)
	if err != nil {
		logger.Error("error when trying to prepare clockOut employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(employee.ClockOut, employee.EmployeeID)
	if err != nil {
		logger.Error("error when trying to clockOut employee", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (employee *TimeCard) EmployeeTotalTime() *errors.RestErr {
	stmt, err := Client.Prepare(queryTotalTime)
	if err != nil {
		logger.Error("error when trying to prepare totalTime employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(employee.TotalTime, employee.EmployeeID)
	if err != nil {
		logger.Error("error when trying to insert total time of employee", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (employee *TimeCard) GetClockIn() *errors.RestErr {
	stmt, err := Client.Prepare(queryGetClockIn)
	if err != nil {
		logger.Error("error when trying to prepare get employee statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name, &employee.ClockIn); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (employee *TimeCard) GetClockInClockOut() *errors.RestErr {
	stmt, err := Client.Prepare(queryGetClockInClockOut)
	if err != nil {
		logger.Error("error when trying to prepare get employee clock in/clock out statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name, &employee.ClockIn, &employee.ClockOut); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (employee *TimeCard) GetTime() *errors.RestErr {
	stmt, err := Client.Prepare(queryGetEmployeeTimecard)
	if err != nil {
		logger.Error("error when trying to prepare get employee time card statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name, &employee.ClockIn, &employee.ClockOut, &employee.TotalTime); getErr != nil {
		logger.Error("error when trying to get employee time card by ID", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
