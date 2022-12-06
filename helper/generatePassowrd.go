package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type IGenPassword interface {
	GeneratePasswordHash(password string) (string, error)
	ComparePasswordHash(passwordHashed, passwrod string) error
}

var NewGenPassword IGenPassword = &newGenPassword{}

type newGenPassword struct{}

func (g *newGenPassword) GeneratePasswordHash(password string) (string, error) {
	log.Println(password)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return string(passwordHash), err
	}

	return string(passwordHash), nil
}

func (g *newGenPassword) ComparePasswordHash(passwordHashed, passwrod string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(passwrod))
	if err != nil {
		return err
	}

	return nil
}
