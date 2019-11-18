package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/roger-king/tasker/models"
	"github.com/roger-king/tasker/utils"
)

func checkCookie(w http.ResponseWriter, r *http.Request) int {
	cookie, err := r.Cookie("tasker-user")

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return http.StatusUnauthorized
		}
		// For any other type of error, return a bad request status
		return http.StatusBadRequest
	}

	tknStr := cookie.Value

	// Initialize a new instance of `Claims`
	claims := &models.UserClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.TaskerJWTSecret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized
		}
		return http.StatusBadRequest
	}

	if !tkn.Valid {
		return http.StatusUnauthorized
	}
	return http.StatusOK
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusCode := checkCookie(w, r)

		if statusCode != 200 {
			http.Error(w, "unauthorized request", statusCode)
			return
		}

		next.ServeHTTP(w, r)
	})
}
