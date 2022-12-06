package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type IService interface {
	GenerateToken(jwtInput JwtInput) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtInput struct {
	UserID int
	Email  string
	Role   string
}

var SECRETE_KEY = "rahasia"

type jwtService struct{}

var NewJwtService IService = &jwtService{}

func (s *jwtService) GenerateToken(jwtInput JwtInput) (string, error) {
	// ttl := 20 * time.Minute

	claim := jwt.MapClaims{
		"id":    jwtInput.UserID,
		"email": jwtInput.Email,
		"role":  jwtInput.Role,
		// "exp":   time.Now().UTC().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(SECRETE_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRETE_KEY), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
