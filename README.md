# timeCardSimple
Simple version of the TimeCard application. This application allows users to add employees to a database, and allows employees to clock in and clock out. Total time can also be retrieved.

## Launch Application
To launch the application, first make sure you have Golang installed locally. Configure env variables correctly for database init. Next set env variable "is_testing" to "false". Then run:

```
go run main.go
```

This starts the application locally and allows access to the Swagger UI at :8080/swaggerui/

## REST API Calls

1. POST /employees takes a payload that contains the employee's name and DoB and generates an employeeID for that employee
2. GET /employees returns all employees with name, employeeID, and DoB. 
3. GET /employees/{id} returns name, DoB, and employeeID associated to that specific employee. 
4. DELETE /employees/{id} deletes the employee's information
5. POST /employees/ClockIn/{id} takes the employeeID as the parameter and returns the time the employee clocked in UTC
6. POST /employees/ClockOut/{id} takes the employeeID and returns the clock in and clock out time
7. GET /employees/TotalTime/{id} returns employeeID, name, clock in and clock out time, as well as the total time the employee worked.

## Tests

To run tests, first make sure env variable "is_testing" is set to "true", then run this command in your terminal in the directory

```
go test -v
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.



