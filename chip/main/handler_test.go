package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestWorkingDayTrue(t *testing.T) {
	req, err := http.NewRequest("GET", "/working_day/2019-01-04", nil)

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"date": "2019-01-04",
	}

	req = mux.SetURLVars(req, vars)

	response := execute(req)
	checkResponse(t, http.StatusOK, true, response)
}

func TestWeekend(t *testing.T) {
	req, err := http.NewRequest("GET", "/working_day/2019-01-05", nil)

	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"date": "2019-01-05",
	}

	req = mux.SetURLVars(req, vars)

	response := execute(req)
	checkResponse(t, http.StatusOK, false, response)
}

func TestHoliday(t *testing.T) {
	req, err := http.NewRequest("GET", "/working_day/2019-01-01", nil)

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"date": "2019-01-01",
	}

	req = mux.SetURLVars(req, vars)

	response := execute(req)
	checkResponse(t, http.StatusOK, false, response)
}

func TestIncorrectDateFormat(t *testing.T) {
	req, err := http.NewRequest("GET", "/working_day/2019-01-", nil)

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"date": "2019-01-",
	}

	req = mux.SetURLVars(req, vars)

	response := execute(req)
	if http.StatusBadRequest != response.Code {
		t.Errorf("response code incorrect, expected %d but got %d\n", http.StatusBadRequest, response.Code)
	}
}

func execute(req *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateWorkingDays)
	handler.ServeHTTP(response, req)

	return response
}

func checkResponse(t *testing.T, expectedCode int, expectedResponse bool, response *httptest.ResponseRecorder) {
	if expectedCode != response.Code {
		t.Errorf("response code incorrect, expected %d but got %d\n", expectedCode, response.Code)
	}
	var responseResult map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseResult)
	if responseResult["WorkingDay"] != expectedResponse {
		t.Errorf("response incorrect, expected (%v) but got %v", expectedResponse, responseResult["id"])
	}
}
