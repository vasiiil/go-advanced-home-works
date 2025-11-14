package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	secret string
}

type JwtData struct {
	SessionId string
	Phone     string
}

type TContextJwtKey string

const (
	ContextJwtKey TContextJwtKey = "contextJwtKey"
)

func New(secret string) *Jwt {
	return &Jwt{
		secret: secret,
	}
}

func (j *Jwt) Create(jwtData *JwtData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": jwtData.SessionId,
		"phone":     jwtData.Phone,
	})
	s, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *Jwt) Parse(tokenString string) (bool, *JwtData) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return false, nil
	}
	claims := token.Claims.(jwt.MapClaims)
	sessionId := claims["sessionId"].(string)
	phone := claims["phone"].(string)
	return token.Valid, &JwtData{
		SessionId: sessionId,
		Phone:     phone,
	}
}
