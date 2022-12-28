package database

import (
	"timeCardSimple/errors"
	"timeCardSimple/logger"
)

const (
	//Should probably have employees, then ClockIn/ClockOut as seperate Structs
	queryInsertEmployee      = "INSERT INTO timecard(name, dob) VALUES(?, ?);"
	queryGetEmployees        = "SELECT employeeID, name, dob FROM timecard;"
	queryGetEmployee         = "SELECT employeeID, name, dob FROM timecard WHERE employeeID=?;"
	queryDeleteEmployee      = "DELETE FROM timecard WHERE employeeID=?;"
	queryClockIn             = "UPDATE timecard SET clockIn=? WHERE employeeID=?;"
	queryClockOut            = "UPDATE timecard SET clockOut=? WHERE employeeID=?;"
	queryTotalTime           = "UPDATE timecard SET totalTime=? WHERE employeeID=?;"
	queryGetEmployeeTimecard = "SELECT employeeID, name, clockIn, clockOut, totalTime WHERE employeeID=?"
)

func (employee *Employee) Save() *errors.RestErr {
	stmt, err := Client.Prepare(queryInsertEmployee)
	if err != nil {
		logger.Error("error when trying to prepare save employee statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if employee.Name == "" || employee.DateOfBirth == "" {
		return errors.NewBadRequestError("Name or DoB cannot be nil")
	}

	insertResult, saveErr := stmt.Exec(employee.Name, employee.DateOfBirth)
	if saveErr != nil {
		logger.Error("error when trying to save employee", saveErr)
		return errors.NewInternalServerError("database error")
	}

	employeeId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get last insert ID after creating New Employee", err)
		return errors.NewInternalServerError("database error")
	}

	employee.EmployeeID = employeeId
	return nil

}

func (employees *Employee) GetAll() ([]Employee, *errors.RestErr) {
	stmt, err := Client.Prepare(queryGetEmployees)
	if err != nil {
		logger.Error("error when trying to prepare get employees statement", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	res, err := stmt.Query()
	if err != nil {
		logger.Error("error when trying to search rows for employees", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer res.Close()

	results := make([]Employee, 0)
	for res.Next() {
		var employee Employee
		if err := res.Scan(&employee.EmployeeID, &employee.Name, &employee.DateOfBirth); err != nil {
			logger.Error("error when trying to scan employee row in employee struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, employee)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("No employees exist")
	}

	return results, nil
}

func (employee *Employee) Get() *errors.RestErr {
	stmt, err := Client.Prepare(queryGetEmployee)
	if err != nil {
		logger.Error("error when trying to prepare get employee statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name, &employee.DateOfBirth); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (employee *Employee) Delete() *errors.RestErr {
	stmt, err := Client.Prepare(queryDeleteEmployee)
	if err != nil {
		logger.Error("error when trying to prepare delete employee statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(employee.EmployeeID); err != nil {
		logger.Error("error when trying to get user from database", err)
		return errors.NewNotFoundError("Employee does not exist")
	}

	return nil
}

func (employee *Employee) EmployeeClockIn() *errors.RestErr {
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

func (employee *Employee) EmployeeClockOut() *errors.RestErr {
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

func (employee *Employee) EmployeeTotalTime() *errors.RestErr {
	stmt, err := Client.Prepare(queryClockOut)
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

func (employee *Employee) GetTime() *errors.RestErr {
	stmt, err := Client.Prepare(queryGetEmployee)
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
