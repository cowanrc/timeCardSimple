package api

import (
	"encoding/json"
	"net/http"
	"time"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/webapp/api/dto"
)

func (a *API) ClockIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	employee, ok := ctx.Value(employeeContextKey{}).(*employee.Employee)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
	}

	now := time.Now()
	timecard, err := a.TimecardSVC.ClockIn(ctx, employee.ID(), now)
	if err != nil {
		http.Error(w, "error clocking in", http.StatusBadRequest)
		return
	}

	transfer, err := dto.Transform(timecard)
	if err != nil {
		http.Error(w, "error transforming timecard", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(transfer); err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}
