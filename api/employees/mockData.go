package employees

import "encoding/json"

type mockEmployeeResponse struct {
	Name        string `json:"name"`
	EmployeeID  int64  `json:"employeeID"`
	DateCreated string `json:"dateCreated"`
}

var mockEmployeeString = `{"name": "John Smith","employeeID": 1,"dateCreated": "01/01/2000"}`

func mockEmployee(in string) mockEmployeeResponse {
	bytes := []byte(in)
	var m mockEmployeeResponse
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		panic(err)
	}
	return m
}
