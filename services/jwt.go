package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/roger-king/tasker/models"
	"github.com/roger-king/tasker/utils"
	"github.com/sirupsen/logrus"
)

func GenerateJWTToken(user *models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(120 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &models.UserClaims{
		User: models.User{
			Bio:       user.Bio,
			Email:     user.Email,
			GitHubURL: user.GitHubURL,
			Name:      user.Name,
			UserName:  user.UserName,
		},
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	logrus.Info(claims.User.UserName)

	jwtKey := []byte(utils.TaskerJWTSecret)

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, expirationTime, err
}
