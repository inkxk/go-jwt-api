package model

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaim struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
