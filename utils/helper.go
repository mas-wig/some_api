package utils

import (
	"encoding/base64"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}

func Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ToDocument(value any) (doc *bson.D, err error) {
	data, err := bson.Marshal(value)
	if err != nil {
		log.Fatal("could not marshal document value : %w", err)
	}
	err = bson.Unmarshal(data, &doc)
	return
}
