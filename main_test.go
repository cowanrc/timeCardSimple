package main

import (
	"bytes"
	"fmt"
	"os"

	"net/http"
	"testing"
)

var client = http.DefaultClient

var name1 = "John Smith"

var link = "http://localhost:8080/employees"

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

	checkResponseCode(t, http.StatusOK, response.StatusCode)
}

func TestGetEmployee(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, link+"/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	// body, _ := ioutil.ReadAll(response.Body)
	// bodyString := string(body)

	// if !strings.Contains(bodyString, `"Name":"John Smith","employeeID":1,"DoB":"1/1/2000"`) {
	// 	t.FailNow()
	// }
}

func TestDeleteEmployee(t *testing.T) {

	// req, _ := http.NewRequest(http.MethodGet, link+"/1", nil)
	// response := executeRequest(req)

	// checkResponseCode(t, http.StatusOK, response.StatusCode)

	req, _ := http.NewRequest(http.MethodDelete, link+"/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	// req, _ = http.NewRequest(http.MethodGet, link+"/1", nil)
	// response = executeRequest(req)

	// checkResponseCode(t, http.StatusNotFound, response.StatusCode)

}

// func TestBadDelete(t *testing.T) {
// 	req, _ := http.NewRequest(http.MethodDelete, link+"/yellow", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusBadRequest, response.StatusCode)
// }

func TestClockIn(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPut, link+"/ClockIn/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	// body, _ := ioutil.ReadAll(response.Body)
	// bodyString := string(body)

	// if !strings.Contains(bodyString, "clockIn") {
	// 	t.FailNow()
	// }
}

// func TestClockInTwiceError(t *testing.T) {
// 	req, _ := http.NewRequest(http.MethodPost, link+"/ClockIn/1", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusBadRequest, response.StatusCode)

// 	body, _ := ioutil.ReadAll(response.Body)
// 	bodyString := string(body)

// 	if !strings.Contains(bodyString, "User cannot clock in multiple times before clocking out once.") {
// 		t.FailNow()
// 	}
// }

// func TestClockInNotFound(t *testing.T) {

// 	req, _ := http.NewRequest(http.MethodPost, link+"/ClockIn/3", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusNotFound, response.StatusCode)
// 	body, _ := ioutil.ReadAll(response.Body)
// 	bodyString := string(body)

// 	if !strings.Contains(bodyString, "Either this employee has been removed or has yet to be added") {
// 		t.FailNow()
// 	}
// }

func TestClockOut(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPut, link+"/ClockOut/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	// body, _ := ioutil.ReadAll(response.Body)
	// bodyString := string(body)

	// if !strings.Contains(bodyString, "clockOut") {
	// 	t.FailNow()
	// }
}

// func TestClockOutTwiceError(t *testing.T) {
// 	req, _ := http.NewRequest(http.MethodPost, link+"/ClockOut/1", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusBadRequest, response.StatusCode)

// 	body, _ := ioutil.ReadAll(response.Body)
// 	bodyString := string(body)

// 	if !strings.Contains(bodyString, "You cannot clock out multiple times. Without Clocking in again first") {
// 		t.FailNow()
// 	}
// }

func TestTotalTime(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, link+"/TotalTime/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.StatusCode)
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
