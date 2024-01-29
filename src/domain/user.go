package domain

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/soat-2023-gp14/soat23-gp14_fiap-lambda-application/models"
	"github.com/soat-2023-gp14/soat23-gp14_fiap-lambda-application/services"
)

type Users struct {
	provider services.Auth
}

var (
	firstDigitTable    = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	secondDigitTable   = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	ErrInvalidDocument = errors.New("document informed is invalid")
)

func NewUsersDomain(p services.Auth) *Users {
	return &Users{
		provider: p,
	}
}

func (u *Users) CreateUser(ctx context.Context, form models.UserForm) error {
	cleanDocument := cleanDocument(form.Document)
	isValid := validateDocument(cleanDocument)
	if !isValid {
		return ErrInvalidDocument
	}
	err := u.provider.SignUp(ctx, models.UserForm{
		Name:     cleanDocument,
		Document: cleanDocument,
		Password: form.Password,
		Email:    form.Email,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) Login(ctx context.Context, form models.UserLogin) (string, error) {
	cleanDocument := cleanDocument(form.Username)
	isValid := validateDocument(cleanDocument)
	if !isValid {
		return "", ErrInvalidDocument
	}
	accessToken, err := u.provider.Login(ctx, form)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func cleanDocument(document string) string {
	regex := regexp.MustCompile("[^0-9]+")
	clean := regex.ReplaceAllString(document, "")
	return clean
}

func validateDocument(document string) bool {
	if len(document) != 11 {
		return false
	}

	firstPart := document[0:9]
	sum := sumDigit(firstPart, firstDigitTable)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)
	dsum := sumDigit(secondPart, secondDigitTable)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == document
}

func sumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}
	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}
	return sum
}
