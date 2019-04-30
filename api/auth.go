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

func AuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isValidToken(r) {
			w.WriteHeader(http.StatusUnauthorized)
			writeResponse(w, nil)
			return
		}

		next(w, r)
	})
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
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

	if req.Password != password {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, expires, err := newToken()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := struct {
		Token   string    `json:"token"`
		Expires time.Time `json:"expires"`
	}{token, expires}
	writeResponse(w, &resp)
}

func newToken() (string, time.Time, error) {
	exp := time.Now().Add(tokenExpirey)
	now := time.Now()
	claims := jwt.StandardClaims{
		ExpiresAt: exp.Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		logrus.WithField("err", err).Error("Error generating token")
		return "", time.Time{}, err
	}

	return tokenString, exp, nil
}

func isValidToken(r *http.Request) bool {
	bearer := strings.Split(r.Header.Get(headerAuth), " ")

	if len(bearer) != 2 || bearer[0] != headerBearer {
		return false
	}

	token := bearer[1]
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected token signing method: %v", token.Header["alg"])
		}

		return verifyKey, nil
	})
	return err == nil
}
