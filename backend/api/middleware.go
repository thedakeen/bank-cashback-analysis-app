package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) requireNoXAuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("X-Auth")
		claims := &AppClaims{}
		if tokenString != "" {
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("s7Ndh+pPznbHbS*+9Pk8qGWhTzbpa@tw"), nil
			})

			if err == nil && token.Valid {
				app.forbiddenRequestResponse(w, r, err)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// check if user authenticated (e.g. access for cart)
func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("X-Auth")
		if tokenString == "" {
			app.unauthorizedResponse(w, r)
			return
		}

		claims := &AppClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("s7Ndh+pPznbHbS*+9Pk8qGWhTzbpa@tw"), nil
		})

		if err != nil || !token.Valid {
			app.forbiddenRequestResponse(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func MiddleCORS(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter,
		r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r, ps)
	}
}
