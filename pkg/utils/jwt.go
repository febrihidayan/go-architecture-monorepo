package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type StandardJwtClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
	Name      string `json:"name,omitempty"`
}

func DecodeJwtToken(token string) (claims *StandardJwtClaims, err error) {

	parts := strings.Split(token, ".")

	if len(token) < 3 {
		return nil, errors.New("invalid token.")
	}

	decoded, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(bytes.NewBuffer(decoded)).Decode(&claims); err != nil {
		return nil, err
	}

	return
}
