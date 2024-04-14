package main

import (
	"bank-cashback-analysis/backend/pkg/models"
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"net/http"
	"regexp"
)

func (app *application) signupEmail(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, validation.Length(5, 100), is.Email),
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	err = app.users.CheckEmail(req.Email)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": "Email already in use"})
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.otps.SignUpEmail(req.Email)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"email": req.Email}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) signupCode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Code  string `json:"code"`
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	isValid, err := app.otps.SignUpConfirmCode(req.Email, req.Code)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid or expired code"})
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"email": req.Email}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) signupFinish(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Address  string `json:"address"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, validation.Length(5, 100), is.Email),
		validation.Field(&req.Name, validation.Required, validation.Length(2, 25), validation.Match(regexp.MustCompile("^[a-zA-Z]+$")).Error("letters only")),
		validation.Field(&req.Password, validation.Required, validation.Length(5, 30)),
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	verified, err := app.otps.IsEmailVerified(req.Email)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !verified {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email not verified"})
		return
	}

	err = app.users.SignUpComplete(req.Email, req.Name, req.Surname, req.Phone, req.Address, req.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": req}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Signup successful"})
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	err = validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, validation.Length(5, 100), is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(5, 30)),
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	userId, userRole, err := app.users.Authenticate(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect email or password"})
		return
	}

	jwt, err := app.generateJWTsignIn(userId, req.Email, userRole)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"token": jwt}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
