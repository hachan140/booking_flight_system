package jwt

import (
	"booking-flight-system/config"
	"booking-flight-system/ent"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service struct {
	secretKey []byte
}

func NewJWTService(conf config.JWTConfig) *Service {
	return &Service{
		secretKey: []byte(conf.SecretKey),
	}
}

type Claim struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

func (s *Service) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claim)
	if ok {
		return claims.Email, nil
	}
	return "", errors.New("error when extract token")
}

func (s *Service) GenerateToken(user ent.Member) (string, time.Time, error) {
	expiredAt := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claim{})
	claims := token.Claims.(*Claim)
	claims.Email = user.Email
	claims.ExpiresAt = jwt.NewNumericDate(expiredAt)

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", time.Time{}, err
	}
	return tokenString, expiredAt, nil
}
