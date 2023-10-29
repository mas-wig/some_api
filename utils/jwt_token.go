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

func ValidateToken(token string, publicKey string) (interface{}, error) {
	decodePublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode public key %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(decodePublicKey))
	if err != nil {
		return nil, fmt.Errorf("could not parse public key : %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method : %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate : %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate : invalid token")
	}
	return claims["sub"], nil
}
