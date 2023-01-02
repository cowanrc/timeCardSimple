package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"net/http"
	"testing"
)

var client = http.DefaultClient

var (
	link = "http://localhost:8080/employees"

	name1 = "John Smith"

	expectedEmployeeResponse = `{"name":"John Smith","employeeID":1,"dateCreated":"01/01/2000"}`

	expectedClockInResponse   = `{"name":"John Smith","employeeID":1,"clockIn":"Sun Jan  2 01:00:00 UTC 2023"}`
	expectedClockOutResponse  = `{"name":"John Smith","employeeID":1,"clockIn":"Sun Jan  2 01:00:00 UTC 2023","clockOut":"Sun Jan  2 09:00:00 UTC 2023"}`
	expectedTotalTimeResponse = `{"name":"John Smith","employeeID":1,"clockIn":"Sun Jan  2 01:00:00 UTC 2023","clockOut":"Sun Jan  2 09:00:00 UTC 2023","totalTime":"08:00:00"}`
)

func TestMain(m *testing.M) {
	go main()
	code := m.Run()
	os.Exit(code)
}

func TestCreateEmployee(t *testing.T) {
	payload := []byte(`{"Name": "` + name1 + `"}`)
	req, _ := http.NewRequest(http.MethodPost, link, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	bodyString := string(body)

	if !strings.Contains(bodyString, expectedEmployeeResponse) {
		t.FailNow()
	}

}

func TestGetEmployee(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, link+"/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	bodyString := string(body)

	if !strings.Contains(bodyString, expectedEmployeeResponse) {
		t.FailNow()
	}
}

func TestDeleteEmployee(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, link+"/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	req, _ = http.NewRequest(http.MethodDelete, link+"/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)
}

func TestClockIn(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPut, link+"/ClockIn/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	bodyString := string(body)

	if !strings.Contains(bodyString, expectedClockInResponse) {
		t.FailNow()
	}
}

func TestClockOut(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPut, link+"/ClockOut/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	bodyString := string(body)

	if !strings.Contains(bodyString, expectedClockOutResponse) {
		t.FailNow()
	}
}

func TestTotalTime(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, link+"/TotalTime/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	bodyString := string(body)

	if !strings.Contains(bodyString, expectedTotalTimeResponse) {
		t.FailNow()
	}
}

func executeRequest(req *http.Request) *http.Response {
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error has occured: ", err)
	}

	return resp
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
