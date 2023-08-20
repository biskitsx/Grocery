package service

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(username string) string
	ValidateToken(encodedToken string) (*jwt.Token, error) {
}

type jwtService struct {
	secretKey string
}

type authCustomClaims struct {
	Username string `json:"name"`
	jwt.StandardClaims
}

func NewJwtService() JwtService {
	return &jwtService{
		secretKey: "eieiei",
	}
}

func (service *jwtService) GenerateToken(username string) string {
	claims := &authCustomClaims{
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}
