package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	UserID uint
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(jwtData JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": jwtData.UserID,
	})

	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *JWT) Parse(token string) (*JWTData, bool) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, false
	}

	userID := t.Claims.(jwt.MapClaims)["userID"].(float64)

	return &JWTData{UserID: uint(userID)}, t.Valid
}
