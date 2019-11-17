package models

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	User
	jwt.StandardClaims
}
