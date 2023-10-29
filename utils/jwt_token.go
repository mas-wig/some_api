package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(ttl time.Duration, payload interface{}, priviteKey string) (string, error) {
	decodePrivateKey, err := base64.StdEncoding.DecodeString(priviteKey)
	if err != nil {
		return "", fmt.Errorf("could not decode private key %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodePrivateKey)
	if err != nil {
		return "", fmt.Errorf("error => parse key : %w", err)
	}

	var (
		now    = time.Now().UTC()
		claims = make(jwt.MapClaims)
	)

	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error => sign token : %w", err)
	}
	return token, nil
}
