package timeCard

import "encoding/json"

type mockTimeCardResponse struct {
	Name       string `json:"name"`
	EmployeeID int64  `json:"employeeID"`
	ClockIn    string `json:"clockIn,omitempty"`
	ClockOut   string `json:"clockOut,omitempty"`
	TotalTime  string `json:"totalTime,omitempty"`
}

var (
	mockClockInResp   = `{"name": "John Smith","employeeID": 1,"clockIn": "Sun Jan  2 01:00:00 UTC 2023"}`
	mockClockOutResp  = `{"name": "John Smith","employeeID": 1,"clockIn": "Sun Jan  2 01:00:00 UTC 2023","clockOut": "Sun Jan  2 09:00:00 UTC 2023"}`
	mockTotalTimeResp = `{"name": "John Smith","employeeID": 1,"clockIn": "Sun Jan  2 01:00:00 UTC 2023","clockOut": "Sun Jan  2 09:00:00 UTC 2023", "totalTime": "08:00:00"}`
)

func mockClockIn(in string) mockTimeCardResponse {
	bytes := []byte(in)
	var m mockTimeCardResponse
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func mockClockOut(in string) mockTimeCardResponse {
	bytes := []byte(in)
	var m mockTimeCardResponse
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func mockTotalTime(in string) mockTimeCardResponse {
	bytes := []byte(in)
	var m mockTimeCardResponse
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		panic(err)
	}
	return m
}
