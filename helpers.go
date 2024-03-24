package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

const fileName = "private.pem"

func GenerateToken(email string, userId int64) (string, error) {
	privKey, err := readPrivateKeyFromFile(fileName)
	if err != nil {
		privKey, err = generateKey()
		if err != nil {
			return "", err
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
	})
	return token.SignedString(privKey)
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		privateKey, err := readPrivateKeyFromFile(fileName)
		if err != nil {
			return nil, err
		}
		return extractECDSAPublicKey(privateKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userIdInterface, ok := claims["userId"]
	if !ok {
		return 0, errors.New("userId claim not found in token")
	}

	userId, ok := userIdInterface.(float64)
	if !ok {
		return 0, errors.New("userId claim is not an int64 type")
	}

	return int64(userId), nil
}

func generateKey() (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	err = savePrivateKeyToFile(privateKey, fileName)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func savePrivateKeyToFile(privateKey *ecdsa.PrivateKey, filename string) error {
	keyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}

	keyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	if err := pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes}); err != nil {
		return err
	}
	return nil
}

func readPrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	keyFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyFile)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func extractECDSAPublicKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &privateKey.PublicKey
}
