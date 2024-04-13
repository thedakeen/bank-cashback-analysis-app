package main

import "net/http"

func (app *application) addCard(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CardNumber string `json:"card_number"`
		CardType   string `json:"card_type"`
		BankName   string `json:"bank_name"`
	}

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.cards.SetCard(req.CardNumber, req.CardType, req.BankName)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"card": req}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}