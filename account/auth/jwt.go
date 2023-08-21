package auth

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
	"github.com/golang-jwt/jwt/v4"
)

// CreateToken takes some claims and a private key (either rsa or ec) and returns a signed json web token
func CreateToken(claims map[string]interface{}, kid string, key interface{}) (string, error) {
	switch k := key.(type) {
	case *rsa.PrivateKey:
		{
			token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(claims))
			token.Header["kid"] = kid
			return token.SignedString(k)
		}
	case *ecdsa.PrivateKey:
		{
			token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims(claims))
			token.Header["kid"] = kid
			return token.SignedString(k)
		}
	case []byte:
		{
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
			token.Header["kid"] = kid
			return token.SignedString(k)
		}
	}
	return "", errors.New("invalid private key")
}

func Token(claims jwt.MapClaims, k *model.Key) (*string, error) {
	privBlock, _ := pem.Decode([]byte(k.Certificate))
	if privBlock == nil {
		return nil, fmt.Errorf("error decoding private key PEM block")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}

	//generate token
	token, err := CreateToken(claims, k.ID.Hex(), privKey)
	if err != nil {
		return nil, fmt.Errorf("token not found")
	}

	return &token, nil
}
