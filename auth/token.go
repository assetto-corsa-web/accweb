package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/emvi/logbuch"
	"time"
)

const (
	headerBearer = "Bearer"
	tokenExpirey = time.Hour * 6
)

type TokenClaims struct {
	jwt.StandardClaims

	IsAdmin bool
	IsMod   bool
	IsRO    bool
}

// NewToken returns a new JWT and expiration time for given settings.
func NewToken(isAdmin, isMod, isRO bool) (string, time.Time, error) {
	exp := time.Now().Add(tokenExpirey)
	now := time.Now()
	claims := TokenClaims{jwt.StandardClaims{
		ExpiresAt: exp.Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
	},
		isAdmin,
		isMod,
		isRO,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		logbuch.Error("Error generating token", logbuch.Fields{"err": err})
		return "", time.Time{}, err
	}

	return tokenString, exp, nil
}

func isValidToken(tokenString string) *TokenClaims {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected token signing method: %v", token.Header["alg"])
		}

		return verifyKey, nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	claims, ok := token.Claims.(*TokenClaims)

	if ok {
		return claims
	}

	return nil
}
