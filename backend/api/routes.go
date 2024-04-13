package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	app.authenticate(router)

	router.HandlerFunc(http.MethodGet, "/v1/cashbacks", app.authenticate(app.GetAllCashBacks))
	router.HandlerFunc(http.MethodPost, "/v1/signup/email", app.signupEmail)
	router.HandlerFunc(http.MethodPost, "/v1/signup/code", app.signupCode)
	router.HandlerFunc(http.MethodPost, "/v1/signup", app.signupFinish)

	router.HandlerFunc(http.MethodPost, "/v1/login", app.login)

	return router
}
