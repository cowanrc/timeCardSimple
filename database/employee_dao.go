package database

import (
	"timeCardSimple/errors"
	"timeCardSimple/logger"
)

const (
	queryInsertEmployee = "INSERT INTO timecard(name, dateCreated) VALUES(?, ?);"
	queryGetEmployees   = "SELECT employeeID, name, dateCreated FROM timecard;"
	queryGetEmployee    = "SELECT employeeID, name, dateCreated FROM timecard WHERE employeeID=?;"
	queryDeleteEmployee = "DELETE FROM timecard WHERE employeeID=?;"
)

func Save(employee *Employee) *errors.RestErr {
	stmt, err := Client.Prepare(queryInsertEmployee)
	if err != nil {
		logger.Error("error when trying to prepare save employee statement", err)
		return errors.NewInternalServerError("error when trying to save employee")
	}
	defer stmt.Close()

	if employee.Name == "" {
		return errors.NewBadRequestError("employee name cannot be left empty")
	}

	insertResult, saveErr := stmt.Exec(employee.Name, employee.DateCreated)
	if saveErr != nil {
		logger.Error("error when trying to save employee", saveErr)
		return errors.NewInternalServerError("error when trying to save employee")
	}

	employeeId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get last insert ID after creating New Employee", err)
		return errors.NewInternalServerError("error when trying to save employee")
	}

	employee.EmployeeID = employeeId
	return nil

}

func GetAll() ([]Employee, *errors.RestErr) {
	stmt, err := Client.Prepare(queryGetEmployees)
	if err != nil {
		logger.Error("error when trying to prepare get employees statement", err)
		return nil, errors.NewInternalServerError("error when trying to get users")
	}

	defer stmt.Close()

	res, err := stmt.Query()
	if err != nil {
		logger.Error("error when trying to search rows for employees", err)
		return nil, errors.NewInternalServerError("error when trying to get employees")
	}

	defer res.Close()

	results := make([]Employee, 0)
	for res.Next() {
		var employee Employee
		if err := res.Scan(&employee.EmployeeID, &employee.Name, &employee.DateCreated); err != nil {
			logger.Error("error when trying to scan employee row in employee struct", err)
			return nil, errors.NewInternalServerError("error when trying to get employees")
		}
		results = append(results, employee)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("No employees exist")
	}

	return results, nil
}

func Get(employee *Employee) *errors.RestErr {
	stmt, err := Client.Prepare(queryGetEmployee)
	if err != nil {
		logger.Error("error when trying to prepare get employee statement", err)
		return errors.NewInternalServerError("error when trying to get employees")
	}

	defer stmt.Close()

	result := stmt.QueryRow(employee.EmployeeID)
	if getErr := result.Scan(&employee.EmployeeID, &employee.Name, &employee.DateCreated); getErr != nil {
		logger.Error("error when trying to get employee by ID", getErr)
		return errors.NewNotFoundError("employee might not exist in the system")
	}

	return nil
}

func Delete(employee *Employee) *errors.RestErr {
	stmt, err := Client.Prepare(queryDeleteEmployee)
	if err != nil {
		logger.Error("error when trying to prepare delete employee statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(employee.EmployeeID); err != nil {
		logger.Error("error when trying to get user from database", err)
		return errors.NewNotFoundError("employee might not exist in the system")
	}

	return nil
}
