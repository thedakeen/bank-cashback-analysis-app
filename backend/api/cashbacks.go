package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) getAllCashBacks(w http.ResponseWriter, r *http.Request) {
	p, err := app.promos.GetAllPromos()
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
