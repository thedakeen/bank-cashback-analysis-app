package main

import (
	"net/http"
)

func (app *application) getAllCashBacks(w http.ResponseWriter, r *http.Request) {
	p, err := app.promos.GetAllPromos()
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"promos": p}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
