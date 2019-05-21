package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/rickar/cal"
)

var createHolidaysOnce sync.Once
var calendar = cal.NewCalendar()

const timeLayout = "2006-01-02"

func CalculateWorkingDays(w http.ResponseWriter, r *http.Request) {
	CreateHolidays(calendar)
	vars := mux.Vars(r)
	dateInput := vars["date"]
	date, err := time.Parse(timeLayout, dateInput)

	if err != nil {
		http.Error(w, "wrong date format", http.StatusBadRequest)
		return
	}
	resultOutput := CalculateWorkDays(date, calendar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultOutput)

}

func CreateHolidays(calendar *cal.Calendar) {
	createHolidaysOnce.Do(func() {
		calendar.AddHoliday(
			cal.GBNewYear,
			cal.GBGoodFriday,
			cal.GBEasterMonday,
			cal.GBEarlyMay,
			cal.GBSpringHoliday,
			cal.GBSummerHoliday,
			cal.GBChristmasDay,
			cal.GBBoxingDay)
	})
}

func CalculateWorkDays(date time.Time, calendar *cal.Calendar) WorkDay {
	workDay := WorkDay{WorkingDay: false}
	if calendar.IsWorkday(date) {
		workDay.WorkingDay = true
	}
	return workDay
}
