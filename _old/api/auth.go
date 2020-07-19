package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const (
	headerAuth   = "Authorization"
	headerBearer = "Bearer"
	tokenExpirey = time.Hour * 6
)

type TokenClaims struct {
	jwt.StandardClaims

	IsAdmin bool
	IsMod   bool
	IsRO    bool
}

type HttpHandler func(http.ResponseWriter, *http.Request, *TokenClaims)

func AuthMiddleware(next HttpHandler, requiresAdmin, requiresMod bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getTokenFromHeader(r)
		claims := isValidToken(token, requiresAdmin, requiresMod)

		if claims == nil {
			w.WriteHeader(http.StatusUnauthorized)
			writeResponse(w, nil)
			return
		}

		next(w, r, claims)
	})
}

func TokenHandler(w http.ResponseWriter, r *http.Request, claims *TokenClaims) {
	writeResponse(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Password string `json:"password"`
	}{}

	if err := decodeJSON(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isAdmin := req.Password == adminPassword
	isMod := req.Password == modPassword || isAdmin
	isRO := req.Password == roPassword || isMod

	if !isAdmin && !isMod && !isRO {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, expires, err := newToken(isAdmin, isMod, isRO)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := struct {
		Token    string    `json:"token"`
		Expires  time.Time `json:"expires"`
		Admin    bool      `json:"admin"`
		Mod      bool      `json:"mod"`
		ReadOnly bool      `json:"read_only"`
	}{token, expires, isAdmin, isMod, isRO}
	writeResponse(w, &resp)
}

func newToken(isAdmin, isMod, isRO bool) (string, time.Time, error) {
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
		logrus.WithField("err", err).Error("Error generating token")
		return "", time.Time{}, err
	}

	return tokenString, exp, nil
}

func getTokenFromHeader(r *http.Request) string {
	bearer := strings.Split(r.Header.Get(headerAuth), " ")

	if len(bearer) != 2 || bearer[0] != headerBearer {
		return ""
	}

	return bearer[1]
}

func isValidToken(tokenString string, requiresAdmin, requiresMod bool) *TokenClaims {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected token signing method: %v", token.Header["alg"])
		}

		return verifyKey, nil
	})

	if err != nil {
		return nil
	}

	claims, ok := token.Claims.(*TokenClaims)

	if ok && token.Valid {
		if requiresAdmin && !claims.IsAdmin || requiresMod && !claims.IsMod {
			return nil
		}
	} else {
		return nil
	}

	return claims
}
