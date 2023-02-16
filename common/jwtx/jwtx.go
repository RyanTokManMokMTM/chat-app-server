package jwtx

import "github.com/golang-jwt/jwt/v4"

func GetToken(iat int64, expired int64, secretKey string, payload map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + expired
	for v, k := range payload {
		claims[v] = k
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
