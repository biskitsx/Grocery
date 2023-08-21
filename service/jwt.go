package service

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(id uint) string
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

type authCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func NewJwtService() JwtService {
	return &jwtService{
		secretKey: "eieiei",
	}
}

func (service *jwtService) GenerateToken(id uint) string {
	claims := &authCustomClaims{
		ID: id,
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
			return nil, errors.New("invalid token")

		}
		return []byte(service.secretKey), nil
	})
}
