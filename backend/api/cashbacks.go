package main

import (
	"bank-cashback-analysis/backend/pkg/models/mongodb"
	"net/http"
	"strconv"
)

func (app *application) getAllCashBacks(w http.ResponseWriter, r *http.Request) {

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 100
	}

	filters := mongodb.Filters{
		Page:     page,
		PageSize: pageSize,
	}

	p, metadata, err := app.promos.GetAllPromos(filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"promos": p, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
