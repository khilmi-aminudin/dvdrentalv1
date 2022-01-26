package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
)

type JWTService interface {
	GenerateToken(username string, isSignin bool) string
	ValidateToken(tokenstring string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name     string `json:"name"`
	IsSignin bool   `json:"is_signin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    getIssuer(),
	}
}

func (service *jwtService) GenerateToken(username string, isSignin bool) string {
	claims := &jwtCustomClaims{
		username,
		isSignin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	signToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := signToken.SignedString([]byte(service.secretKey))
	helper.PanicIfError(err)

	return token
}

func (service *jwtService) ValidateToken(tokenstring string) (*jwt.Token, error) {
	return jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func getSecretKey() string {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func getIssuer() string {
	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		issuer = "issuer"
	}
	return issuer
}
