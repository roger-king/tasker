package handlers

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/roger-king/tasker/models"
	"github.com/roger-king/tasker/utils"
)

type ContextKey string

func checkCookie(w http.ResponseWriter, r *http.Request) (int, *models.UserClaims) {
	cookie, err := r.Cookie("tasker-user")

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return http.StatusUnauthorized, nil
		}
		// For any other type of error, return a bad request status
		return http.StatusBadRequest, nil
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
			return http.StatusUnauthorized, nil
		}
		return http.StatusBadRequest, nil
	}

	if !tkn.Valid {
		return http.StatusUnauthorized, nil
	}

	// Adding user to context

	return http.StatusOK, claims
}

func getCurrentUser(ctx context.Context) *models.User {
	userCtx := ctx.Value(ContextKey("user"))

	if userCtx != nil {
		user := userCtx.(models.User)
		return &user
	}

	return nil
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusCode, userClaims := checkCookie(w, r)

		if statusCode != 200 {
			http.Error(w, "unauthorized request", statusCode)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKey("user"), userClaims.User)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
